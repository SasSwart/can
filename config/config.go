package config

import (
	"flag"
	"fmt"
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// TODO this may be vague in definition for the sake of it's legibility in use
type Data struct {
	Generator struct {
		ModuleName string

		// BasePackageName represents the
		BasePackageName string
	}

	Template struct {
		// TODO flag that lists available template names?
		// TODO should exit(1) if not set
		Name string

		// Directory should be ./templates/${Name} by default
		Directory string
	}

	// OpenAPIFile represents the path to the yaml OpenAPI 3 file to render
	OpenAPIFile    string
	AbsOpenAPIPath string

	// OutputPath is ./gen by default
	OutputPath string

	// WorkingDirectory is set through calling os.Getwd()
	WorkingDirectory string

	// FilePath is `.` if not set through the `--configFile` flag
	FilePath string
}

func (d Data) init() {

}
func (d Data) Load() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Config.load:: could not determine working directory: %w\n", err)
	}
	var configFilePath *string
	var args *flag.FlagSet

	if errors.DEBUG {
		args = flag.NewFlagSet("can", flag.ContinueOnError)
	} else {
		args = flag.NewFlagSet("can", flag.ExitOnError)
	}
	configFilePath = args.String("configFile", ".", "Specify which config file to use")
	//var configFilePath = args.String("configFile", "", "Specify which config file to use")
	args.BoolVar(&errors.DEBUG, "DEBUG", false, "Enable debug logging")
	_ = args.Parse(os.Args[1:])

	exe, err := os.Readlink("/proc/self/exe")
	if err != nil {
		return fmt.Errorf("setupRenderer:: could not read /proc/self/exe: %w", err)
	}
	fmt.Printf("Using config file \"%s\".\n", *configFilePath)
	viper.SetConfigFile(*configFilePath)

	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("loadConfig:: could not read config file: %w\n", err)
	}

	d.WorkingDirectory = wd
	d.FilePath = viper.ConfigFileUsed()
	d.absOpenAPIPath()

	err = viper.Unmarshal(&d)
	if err != nil {
		return fmt.Errorf("loadConfig:: could not parse config file: %w\n", err)
	}

	return nil
}

// absoluteOpenAPIFile uses the current working directory, resolved config file and the openAPI file that was specified
// in the config file to determine the absolute path to an OpenAPI file. It takes into account that any of these,
// except the working directory could be relative.
func (d Data) absOpenAPIPath() {
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
