package swift

import (
	"os"
	"os/exec"

	"github.com/bitrise-io/go-utils/command"
)

/*
swift build \
	[--build-path <buildPath>] \
	[--configuration <configuration>] \
	[--build-tests]
	[--disable-sandbox]
*/

// BuildCommandModel ...
type BuildCommandModel struct {
	buildPath string

	// Build Settings
	configuration string
	buildTests    bool

	// Options
	disableSandbox bool
}

// NewBuildCommand ...
func NewBuildCommand() *BuildCommandModel {
	return &BuildCommandModel{}
}

// SetBuildPath ...
func (c *BuildCommandModel) SetBuildPath(buildPath string) *BuildCommandModel {
	c.buildPath = buildPath
	return c
}

// SetConfiguration ...
func (c *BuildCommandModel) SetConfiguration(configuration string) *BuildCommandModel {
	c.configuration = configuration
	return c
}

// SetBuildTests ...
func (c *BuildCommandModel) SetBuildTests(buildTests bool) *BuildCommandModel {
	c.buildTests = buildTests
	return c
}

// SetDisableSandbox ...
func (c *BuildCommandModel) SetDisableSandbox(disableSandbox bool) *BuildCommandModel {
	c.disableSandbox = disableSandbox
	return c
}

// PrintableCmd ...
func (c *BuildCommandModel) PrintableCmd() string {
	cmdSlice := c.cmdSlice()
	return command.PrintableCommandArgs(false, cmdSlice)
}

// Command ...
func (c *BuildCommandModel) Command() *command.Model {
	cmdSlice := c.cmdSlice()
	return command.New(cmdSlice[0], cmdSlice[1:]...)
}

// Cmd ...
func (c *BuildCommandModel) Cmd() *exec.Cmd {
	command := c.Command()
	return command.GetCmd()
}

// Run ...
func (c *BuildCommandModel) Run() error {
	command := c.Command()

	command.SetStdout(os.Stdout)
	command.SetStderr(os.Stderr)

	return command.Run()
}

func (c *BuildCommandModel) cmdSlice() []string {
	slice := []string{toolName}
	slice = append(slice, "build")

	if c.configuration != "" {
		slice = append(slice, "--configuration", c.configuration)
	}

	if c.buildPath != "" {
		slice = append(slice, "--build-path", c.buildPath)
	}

	if c.buildTests {
		slice = append(slice, "--build-tests")
	}

	if c.disableSandbox {
		slice = append(slice, "--disable-sandbox")
	}

	return slice
}
