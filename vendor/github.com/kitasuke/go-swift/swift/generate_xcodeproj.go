package swift

import (
	"os"
	"os/exec"

	"github.com/bitrise-io/go-utils/command"
)

/*
swift package generate-xcodeproj
	[--output <output>] \
	[--enable-code-coverage] \
	[--xcconfig-overrides <xcconfig-overrides>] \
*/

// GenerateXcodeprojCommandModel ...
type GenerateXcodeprojCommandModel struct {
	output string

	// Options
	enableCodeCoverage    bool
	xcconfigOverridesPath string
}

// NewGenerateXcodeprojCommand ...
func NewGenerateXcodeprojCommand() *GenerateXcodeprojCommandModel {
	return &GenerateXcodeprojCommandModel{}
}

// SetOutput ...
func (c *GenerateXcodeprojCommandModel) SetOutput(output string) *GenerateXcodeprojCommandModel {
	c.output = output
	return c
}

// SetEnableCodeCoverage ...
func (c *GenerateXcodeprojCommandModel) SetEnableCodeCoverage(enableCodeCoverage bool) *GenerateXcodeprojCommandModel {
	c.enableCodeCoverage = enableCodeCoverage
	return c
}

// SetXcconfigOverridesPath ...
func (c *GenerateXcodeprojCommandModel) SetXcconfigOverridesPath(xcconfigOverridesPath string) *GenerateXcodeprojCommandModel {
	c.xcconfigOverridesPath = xcconfigOverridesPath
	return c
}

// PrintableCmd ...
func (c *GenerateXcodeprojCommandModel) PrintableCmd() string {
	cmdSlice := c.cmdSlice()
	return command.PrintableCommandArgs(false, cmdSlice)
}

// Command ...
func (c *GenerateXcodeprojCommandModel) Command() *command.Model {
	cmdSlice := c.cmdSlice()
	return command.New(cmdSlice[0], cmdSlice[1:]...)
}

// Cmd ...
func (c *GenerateXcodeprojCommandModel) Cmd() *exec.Cmd {
	command := c.Command()
	return command.GetCmd()
}

// Run ...
func (c *GenerateXcodeprojCommandModel) Run() error {
	command := c.Command()

	command.SetStdout(os.Stdout)
	command.SetStderr(os.Stderr)

	return command.Run()
}

func (c *GenerateXcodeprojCommandModel) cmdSlice() []string {
	slice := []string{toolName}
	slice = append(slice, "package")
	slice = append(slice, "generate-xcodeproj")

	if c.output != "" {
		slice = append(slice, "--output", c.output)
	}

	if c.enableCodeCoverage {
		slice = append(slice, "--enable-code-coverage")
	}

	if c.xcconfigOverridesPath != "" {
		slice = append(slice, "--xcconfig-overrides", c.xcconfigOverridesPath)
	}

	return slice
}
