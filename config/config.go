package config

import (
	"flag"
	"fmt"
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Data represents the config data used in the day-to-day running of can
//	TODO this may be vague in definition for the sake of its legibility in use
type Data struct {
	Generator struct {
		ModuleName string

		// BasePackageName represents the
		BasePackageName string
	}

	TemplatesDir string
	Template     struct {
		// TODO flag that lists available template names?
		// TODO should exit(1) if not set
		Name string

		// Directory should be ./templates/${Name} by default
		Directory    string
		AbsDirectory string
	}

	// OpenAPIFile represents the path to the yaml OpenAPI 3 file to render
	OpenAPIFile    string
	AbsOpenAPIPath string

	OutputPath string

	// WorkingDirectory is set through calling os.Getwd()
	WorkingDirectory string

	// FilePath is `.` if not set through the `--configFile` flag
	FilePath string
}

func (d Data) Load() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Config.load:: could not determine working directory: %w\n", err)
	}
	var args *flag.FlagSet

	if errors.DEBUG {
		args = flag.NewFlagSet("can", flag.ContinueOnError)
	} else {
		args = flag.NewFlagSet("can", flag.ExitOnError)
	}

	configFilePath := args.String("configFile", ".", "Specify which config file to use")
	template := args.String("template", "", "Specify which template set to use")
	args.BoolVar(&errors.DEBUG, "DEBUG", false, "Enable debug logging")
	_ = args.Parse(os.Args[1:])

	if template == nil {
		fmt.Printf("template is a required flag\nexiting...")
		os.Exit(1)
	}
	if !d.validTemplateName(*template) {
		fmt.Printf("%s is an invalid flag\nexiting...\n", *template)
		os.Exit(1)
	}

	fmt.Printf("Using config file \"%s\".\n", *configFilePath)
	viper.SetConfigFile(*configFilePath)

	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("loadConfig:: could not read config file: %w\n", err)
	}

	exe, err := os.Readlink("/proc/self/exe")
	if err != nil {
		return fmt.Errorf("setupRenderer:: could not read /proc/self/exe: %w", err)
	}

	d.WorkingDirectory = wd
	d.FilePath = viper.ConfigFileUsed()
	d.TemplatesDir = filepath.Join(filepath.Dir(exe), "templates")
	d.absOpenAPIPaths()
	d.absTemplateDirs()

	err = viper.Unmarshal(&d)
	if err != nil {
		return fmt.Errorf("loadConfig:: could not parse config file: %w\n", err)
	}

	return nil
}
func (d Data) validTemplateName(name string) bool {
	dirs, err := os.ReadDir(d.TemplatesDir)
	if err != nil {
		fmt.Println(fmt.Errorf("could not list directories %w", err))
		return false
	}
	for _, dir := range dirs {
		if dir.Name() == name {
			return true
		}
	}
	return false
}

// absoluteOpenAPIFile uses the current working directory, resolved config file and the openAPI file that was specified
// in the config file to determine the absolute path to an OpenAPI file. It takes into account that any of these,
// except the working directory could be relative.
func (d Data) absOpenAPIPaths() {
	if filepath.IsAbs(d.OpenAPIFile) {
		d.AbsOpenAPIPath = d.OpenAPIFile
	} else {
		if filepath.IsAbs(d.FilePath) {
			d.AbsOpenAPIPath = filepath.Join(
				filepath.Dir(d.FilePath),
				d.OpenAPIFile,
			)
		} else {
			d.AbsOpenAPIPath = filepath.Join(
				// not relative as per above comment
				d.WorkingDirectory,
				filepath.Dir(d.FilePath),
				d.OpenAPIFile,
			)
		}
	}
}
func (d Data) absTemplateDirs() {
	switch true {
	case filepath.IsAbs(d.TemplatesDir):
		d.Template.AbsDirectory = filepath.Join(d.TemplatesDir, d.Template.Name)
	case filepath.IsAbs(d.FilePath):
		d.Template.AbsDirectory = filepath.Join(filepath.Dir(d.FilePath), d.TemplatesDir, d.Template.Name)
	default:
		d.Template.AbsDirectory = filepath.Join(d.WorkingDirectory, filepath.Dir(d.FilePath), d.TemplatesDir, d.Template.Name)
	}
}
