package swift

import "github.com/bitrise-io/go-utils/command"

const (
	toolName = "swift"
)

// CommandModel
type CommandModel interface {
	PrintableCmd() string
	Command() *command.Model
}
