package swift

import (
	"os"
	"os/exec"

	"github.com/bitrise-io/go-utils/command"
)

/*
swift test \
	[--build-path <buildPath>] \
	[--configuration <configuration>] \
	[--skip-build] \
	[--parallel] \
*/

// TestCommandModel ...
type TestCommandModel struct {
	buildPath string

	// Build Settings
	skipBuild bool

	// Options
	isParallel bool
}

// NewTestCommand ...
func NewTestCommand() *TestCommandModel {
	return &TestCommandModel{}
}

// SetBuildPath ...
func (c *TestCommandModel) SetBuildPath(buildPath string) *TestCommandModel {
	c.buildPath = buildPath
	return c
}

// SetSkipBuild ...
func (c *TestCommandModel) SetSkipBuild(skipBuild bool) *TestCommandModel {
	c.skipBuild = skipBuild
	return c
}

// SetParallel ...
func (c *TestCommandModel) SetIsParallel(isParallel bool) *TestCommandModel {
	c.isParallel = isParallel
	return c
}

// PrintableCmd ...
func (c *TestCommandModel) PrintableCmd() string {
	cmdSlice := c.cmdSlice()
	return command.PrintableCommandArgs(false, cmdSlice)
}

// Command ...
func (c *TestCommandModel) Command() *command.Model {
	cmdSlice := c.cmdSlice()
	return command.New(cmdSlice[0], cmdSlice[1:]...)
}

// Cmd ...
func (c *TestCommandModel) Cmd() *exec.Cmd {
	command := c.Command()
	return command.GetCmd()
}

// Run ...
func (c *TestCommandModel) Run() error {
	command := c.Command()

	command.SetStdout(os.Stdout)
	command.SetStderr(os.Stderr)

	return command.Run()
}

func (c *TestCommandModel) cmdSlice() []string {
	slice := []string{toolName}
	slice = append(slice, "test")

	if c.buildPath != "" {
		slice = append(slice, "--build-path", c.buildPath)
	}

	if c.skipBuild {
		slice = append(slice, "--skip-build")
	}

	if c.isParallel {
		slice = append(slice, "--parallel")
	}

	return slice
}
