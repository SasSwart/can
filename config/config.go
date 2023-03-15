package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// SemVer should be updated on any new release!!
const SemVer = "0.0.7"

var (
	Debug          bool
	Dryrun         bool
	VersionFlagSet bool
	// ConfigFilePath is `.` if not set through the `-configFile` flag
	ConfigFilePath string
	// ProcWorkingDir is set through calling os.Getwd()
	ProcWorkingDir string
	ExePath        string

	// for package use only
	outputPathFlag   string
	templateNameFlag string
)

// Data represents the config data used in the day-to-day running of can
//
//	TODO this may be vague in definition for the sake of its legibility in use
type Data struct {
	Template

	// left public due to it's need to be unmarshalled by Data.Load()
	TemplatesDir string `yaml:"templatesDir"`

	// OpenAPIFile represents the path to the yaml OpenAPI 3 file to render
	OpenAPIFile    string
	absOpenAPIPath string

	OutputPath    string
	absOutputPath string
}

// Template's fields are dependant on there being a defined templated name in the CLI arguments or config file.
type Template struct {
	Name       string
	ModuleName string
	// BasePackageName represents the
	BasePackageName string

	absDirectory string
}

func (d *Data) Load() (err error) {
	// Setup config pre-unmarshalling. This assumes we don't change directory before this is executed
	ProcWorkingDir, err = os.Getwd()
	if err != nil {
		return fmt.Errorf("Config.load:: could not determine working directory: %w\n", err)
	}
	ExePath, err = os.Executable()
	if err != nil {
		return fmt.Errorf("Config.load:: could not determine executable path: %w\n", err)
	}

	// flags
	if !flag.Parsed() { // Sorts out buggy tests
		flag.BoolVar(&VersionFlagSet, "version", false, "Print Can version and exit")
		flag.BoolVar(&Debug, "debug", false, "Enable debug logging")
		flag.StringVar(&ConfigFilePath, "configFile", ".", "Specify which config file to use")
		flag.BoolVar(&Dryrun, "dryrun", false,
			"Toggles whether to perform a render without writing to disk."+
				"This works particularly well in combination with -debug")
		flag.StringVar(&templateNameFlag, "template", "", "Specify which template set to use")
		// TODO should we support stdout as an option here?
		flag.StringVar(&outputPathFlag, "outputPath", "", "Specify where to write output to")
		flag.Parse()
	}

	absCfgPath, err := filepath.Abs(ConfigFilePath)
	if err != nil {
		return fmt.Errorf("could not resolve relative config path: %w", err)
	}
	ConfigFilePath = absCfgPath

	if VersionFlagSet {
		fmt.Printf("Can: v%s\n", SemVer)
		os.Exit(0)
	}

	// config load
	if Debug {
		fmt.Printf("[v%s]::Using config file \"%s\".\n", SemVer, ConfigFilePath)
	}
	err = d.setOverridesAndLoadConfig()
	if err != nil {
		return err
	}

	err = d.resolveTemplateConfig()
	if err != nil {
		return err
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
	return nil
}

func (d *Data) GetTemplateFilesDir() (path string) {
	if d.Template.absDirectory != "" {
		return d.Template.absDirectory
	}
	switch true {
	case filepath.IsAbs(d.TemplatesDir):
		d.Template.absDirectory = filepath.Join(d.TemplatesDir, d.Template.Name)
		return d.Template.absDirectory
	case filepath.IsAbs(ConfigFilePath):
		d.Template.absDirectory = filepath.Join(filepath.Dir(ConfigFilePath), d.TemplatesDir, d.Template.Name)
		return d.Template.absDirectory

	// No absolute dir provided. Let's build one
	default:
		d.Template.absDirectory = filepath.Join(ProcWorkingDir, filepath.Dir(ConfigFilePath), d.TemplatesDir, d.Template.Name)
		return d.Template.absDirectory
	}
}

// GetOutputDir is used by the render engine to determine where rendered files will be written to
func (d *Data) GetOutputDir() (path string) {
	if d.absOutputPath != "" {
		return d.absOutputPath
	}
	switch true {
	case filepath.IsAbs(d.OutputPath):
		d.absOutputPath = d.OutputPath
		return d.absOutputPath
	case filepath.IsAbs(ConfigFilePath):
		d.absOutputPath = filepath.Join(
			filepath.Dir(ConfigFilePath),
			d.OutputPath,
		)
		return d.absOutputPath

	// No absolute dir provided. Let's build one
	default:
		d.absOutputPath = filepath.Join(
			ProcWorkingDir,
			filepath.Dir(ConfigFilePath),
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
		if filepath.IsAbs(ConfigFilePath) {
			d.absOpenAPIPath = filepath.Join(
				filepath.Dir(ConfigFilePath),
				d.OpenAPIFile,
			)
			return d.absOpenAPIPath
		} else {
			d.absOpenAPIPath = filepath.Join(
				// not relative as per above comment
				ProcWorkingDir,
				filepath.Dir(ConfigFilePath),
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
	pwdDirs, err := os.ReadDir(ProcWorkingDir)
	if err != nil {
		return err
	}
	exeDirs, err := os.ReadDir(filepath.Dir(ExePath))
	if err != nil {
		return err
	}
	if d.TemplatesDir == "" {
		// First we look in the process directory
		for _, dirEnt := range pwdDirs {
			if dirEnt.Name() == "templates" {
				templateDir := filepath.Join(ProcWorkingDir, "templates")
				d.TemplatesDir = templateDir
				fmt.Printf("No template directory specified, defaulting to %s\n", templateDir)
				break
			}
		}
		// Then we look in the executable directory
		if d.TemplatesDir == "" {
			for _, dirEnt := range exeDirs {
				if dirEnt.Name() == "templates" {
					templateDir := filepath.Join(filepath.Dir(ExePath), "templates")
					d.TemplatesDir = templateDir
					fmt.Printf("No template directory specified, defaulting to %s\n", templateDir)
					break
				}
			}
		}
	}
	if !filepath.IsAbs(d.TemplatesDir) {
		d.TemplatesDir, err = filepath.Abs(d.TemplatesDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Data) setOverridesAndLoadConfig() error {
	viper.SetConfigFile(ConfigFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("setOverridesAndLoadConfig:: could not read config file: %w\n", err)
	}
	// Handle Template name
	err = viper.Unmarshal(&d)
	if err != nil {
		return fmt.Errorf("setOverridesAndLoadConfig:: could not unmarsha; config file: %w\n", err)
	}
	// Handle empty config fields
	if d.Template.Name == "" {
		fmt.Printf("template not set via config file. Expecting field to be set via command line\n")
		if templateNameFlag == "" {
			return fmt.Errorf("template not set via cli flags either\nexiting...")
		}
		d.Template.Name = templateNameFlag
	}
	if d.OutputPath == "" {
		fmt.Printf("outputPath not set via config file. Expecting field to be set via command line\n")
		if outputPathFlag == "" {
			return fmt.Errorf("outputPath not set via cli flags either\nexiting...")
		}
		d.OutputPath = outputPathFlag
	}
	return nil
}
