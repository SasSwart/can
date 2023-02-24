package config

import (
	"flag"
	"fmt"
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// SemVer should be updated on any new release!!
const SemVer = "0.0.6"

// Data represents the config data used in the day-to-day running of can
//
//	TODO this may be vague in definition for the sake of its legibility in use
//	TODO check for redundancy
type Data struct {
	Generator
	Template

	// left public due to it's need to be unmarshalled by Data.Load()
	TemplatesDir string `yaml:"templatesDir"`

	// OpenAPIFile represents the path to the yaml OpenAPI 3 file to render
	OpenAPIFile    string
	absOpenAPIPath string

	OutputPath    string
	absOutputPath string

	// workingDirectory is set through calling os.Getwd()
	workingDirectory string

	// ConfigPath is `.` if not set through the `-configFile` flag
	ConfigPath string
}

type Generator struct {
	ModuleName string

	// BasePackageName represents the
	BasePackageName string
}

// Template is populated based on it's Name variable set as a CLI flag
type Template struct {
	Name string

	// Directory should be ./templates/${Name} by default
	Directory    string
	absDirectory string
}

func (d *Data) Load() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Config.load:: could not determine working directory: %w\n", err)
	}
	var args *flag.FlagSet
	var versionFlagSet bool

	if errors.Debug {
		// TODO verify that this is beneficial
		fmt.Printf("[%s]:: continuing on error...\n", SemVer)
		args = flag.NewFlagSet("can", flag.ContinueOnError)
	} else {
		args = flag.NewFlagSet("can", flag.ExitOnError)
	}

	// flags
	cfgPath := args.String("configFile", ".", "Specify which config file to use")
	templateName := args.String("template", "", "Specify which template set to use")
	args.BoolVar(&errors.Debug, "debug", false, "Enable debug logging")
	args.BoolVar(&versionFlagSet, "version", false, "Print Can version and exit")
	err = args.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	absCfgPath, err := filepath.Abs(*cfgPath)
	if err != nil {
		return fmt.Errorf("could not resolve relative config path: %w", err)
	}

	d.ConfigPath = absCfgPath
	d.Template.Name = *templateName

	if versionFlagSet {
		fmt.Printf("Can Version: %s\n", SemVer)
		os.Exit(0)
	}

	// config load
	if errors.Debug {
		fmt.Printf("[%s]::Using config file \"%s\".\n", SemVer, d.ConfigPath)
	}
	viper.SetConfigFile(d.ConfigPath)

	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("loadConfig:: could not read config file: %w\n", err)
	}

	// Setup config pre-unmarshalling
	d.workingDirectory = wd

	err = viper.Unmarshal(&d)
	if err != nil {
		return fmt.Errorf("loadConfig:: could not parse config file: %w\n", err)
	}

	// This should always happen at the end of this function
	// Handle Templates
	if d.Template.Name == "" {
		fmt.Printf("template is a required flag\nexiting...")
		os.Exit(1)
	}
	if !d.validTemplateName() {
		fmt.Printf("%s does not exist in %s\nexiting...\n", d.Template.Name, d.TemplatesDir)
		fmt.Println("Valid template names are:")
		names, err := d.validTemplates()
		if err != nil {
			return fmt.Errorf("could not read templates: %w", err)
		}
		for _, name := range names {
			fmt.Println(name)
		}
		os.Exit(1)
	}

	// resolve paths
	err = d.resolveTemplateConfig()
	if err != nil {
		return err
	}

	return nil
}

func (d *Data) GetTemplateDir() (path string) {
	if d.Template.absDirectory != "" {
		return d.Template.absDirectory
	}
	switch true {
	case filepath.IsAbs(d.TemplatesDir):
		d.Template.absDirectory = filepath.Join(d.TemplatesDir, d.Template.Name)
		return d.Template.absDirectory
	case filepath.IsAbs(d.ConfigPath):
		d.Template.absDirectory = filepath.Join(filepath.Dir(d.ConfigPath), d.TemplatesDir, d.Template.Name)
		return d.Template.absDirectory
	default:
		d.Template.absDirectory = filepath.Join(d.workingDirectory, filepath.Dir(d.ConfigPath), d.TemplatesDir, d.Template.Name)
		return d.Template.absDirectory
	}
}

// GetOutputFilepath is used by the render engine to determine where rendered files will be written to
func (d *Data) GetOutputFilepath() (path string) {
	if d.absOutputPath != "" {
		return d.absOutputPath
	}
	switch true {
	case filepath.IsAbs(d.OutputPath):
		d.absOutputPath = d.OutputPath
		return d.absOutputPath
	case filepath.IsAbs(d.ConfigPath):
		d.absOutputPath = filepath.Join(
			filepath.Dir(d.ConfigPath),
			d.OutputPath,
		)
		return d.absOutputPath
	default:
		d.absOutputPath = filepath.Join(
			d.workingDirectory,
			filepath.Dir(d.ConfigPath),
			d.OutputPath,
		)
		return d.absOutputPath
	}
}

// GetOpenAPIFilepath uses the current working directory, resolved config file and the openAPI file that was specified
// in the config file to determine the absolute path to an OpenAPI file. It takes into account that any of these,
// except the working directory could be relative. It returns the absolute value on every call by caching the result of
// it's first run and returning that on successive calls
func (d *Data) GetOpenAPIFilepath() (path string) {
	if d.absOpenAPIPath != "" { // we shouldn't have to run below logic multiple times
		return d.absOpenAPIPath
	}
	if filepath.IsAbs(d.OpenAPIFile) {
		d.absOpenAPIPath = d.OpenAPIFile
		return d.absOpenAPIPath
	} else {
		if filepath.IsAbs(d.ConfigPath) {
			d.absOpenAPIPath = filepath.Join(
				filepath.Dir(d.ConfigPath),
				d.OpenAPIFile,
			)
			return d.absOpenAPIPath
		} else {
			d.absOpenAPIPath = filepath.Join(
				// TODO test this
				// not relative as per above comment
				d.workingDirectory,
				filepath.Dir(d.ConfigPath),
				d.OpenAPIFile,
			)
			return d.absOpenAPIPath
		}
	}
}

func (d *Data) validTemplateName() bool {
	dirs, err := d.validTemplates()
	if err != nil {
		fmt.Println(fmt.Errorf("could not list valid templates in %s :: %w", d.TemplatesDir, err))
		return false
	}
	for _, dir := range dirs {
		if dir == d.Template.Name {
			return true
		}
	}
	return false
}

func (d *Data) validTemplates() (templates []string, err error) {
	dirs, err := os.ReadDir(d.TemplatesDir)
	if err != nil {
		fmt.Println(fmt.Errorf("could not list directories %w", err))
		return nil, err
	}
	for _, dir := range dirs {
		templates = append(templates, dir.Name())
	}
	return templates, nil
}
func (d *Data) resolveTemplateConfig() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	if d.TemplatesDir == "" {
		d.TemplatesDir = filepath.Join(filepath.Dir(exe), "templates")
	}
	if d.Template.Directory == "" {
		d.Template.Directory = "./templates/" + d.Template.Name
	}
	return nil
}
