package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/kitasuke/go-swift/swift"
	"github.com/kitasuke/go-swift/utility"
)

const (
	OutputEnvKey             = "output"
	EnableCodeCoverageEnvKey = "enable_code_coverage"
	XcconfigOverridesEnvKey  = "xcconfig_overrides"
)

// ConfigModel ...
type ConfigModel struct {
	// Project Parameters
	output string

	// Build Run Configs
	enableCodeCoverage string
	xcconfigOverrides  string
}

func (configs ConfigModel) print() {
	fmt.Println()

	log.Infof("Project Parameters:")
	log.Printf("- Output: %s", configs.output)

	fmt.Println()
	log.Infof("Build Run Configs:")
	log.Printf("- EnableCodeCoverage: %s", configs.enableCodeCoverage)
	log.Printf("- XcconfigOverrides: %s", configs.xcconfigOverrides)
}

func createConfigsModelFromEnvs() ConfigModel {
	return ConfigModel{
		// Project Parameters
		output: os.Getenv(OutputEnvKey),

		// Configs
		enableCodeCoverage: os.Getenv(EnableCodeCoverageEnvKey),
		xcconfigOverrides:  os.Getenv(XcconfigOverridesEnvKey),
	}
}

func (configs ConfigModel) validate() error {
	if err := validateRequiredInputWithOptions(configs.enableCodeCoverage, EnableCodeCoverageEnvKey, []string{"yes", "no"}); err != nil {
		return err
	}

	return nil
}

//--------------------
// Functions
//--------------------

func validateRequiredInput(value, key string) error {
	if value == "" {
		return fmt.Errorf("Missing required input: %s", key)
	}
	return nil
}

func validateRequiredInputWithOptions(value, key string, options []string) error {
	if err := validateRequiredInput(value, key); err != nil {
		return err
	}

	found := false
	for _, option := range options {
		if option == value {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Invalid input: (%s) value: (%s), valid options: %s", key, value, strings.Join(options, ", "))
	}

	return nil
}

func failf(format string, v ...interface{}) {
	log.Errorf(format, v...)
	os.Exit(1)
}

//--------------------
// Main
//--------------------

func main() {
	configs := createConfigsModelFromEnvs()
	configs.print()
	if err := configs.validate(); err != nil {
		failf("Issue with input: %s", err)
	}

	fmt.Println()
	log.Infof("Other Configs:")

	enableCodeCoverage := configs.enableCodeCoverage == "yes"

	swiftVersion, err := utility.GetSwiftVersion()
	if err != nil {
		failf("Failed to get the version of swift! Error: %s", err)
	}

	log.Printf("* swift_version: %s (%s)", swiftVersion.Version, swiftVersion.Target)

	fmt.Println()

	// setup CommandModel for test
	generateCommandModel := swift.NewGenerateXcodeprojCommand()
	generateCommandModel.SetOutput(configs.output)
	generateCommandModel.SetEnableCodeCoverage(enableCodeCoverage)
	if configs.xcconfigOverrides != "" {
		generateCommandModel.SetXcconfigOverridesPath(configs.xcconfigOverrides)
	}

	log.Infof("$ %s\n", generateCommandModel.PrintableCmd())

	if err := generateCommandModel.Run(); err != nil {
		failf("Generate Xcodeproj failed, error: %s", err)
	}
}
