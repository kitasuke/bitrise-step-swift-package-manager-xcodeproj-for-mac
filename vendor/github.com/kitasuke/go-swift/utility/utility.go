package utility

import (
	"fmt"
	"strings"

	"github.com/bitrise-io/go-utils/command"
	"github.com/kitasuke/go-swift/models"
)

// GetSwiftVersion ...
func GetSwiftVersion() (models.SwiftVersionModel, error) {
	cmd := command.New("swift", "-version")
	outStr, err := cmd.RunAndReturnTrimmedCombinedOutput()
	if err != nil {
		return models.SwiftVersionModel{}, fmt.Errorf("swift -version failed. err: %s, detail: %s", err, outStr)
	}
	return getSwiftVersionFromSwiftOutput(outStr)
}

func getSwiftVersionFromSwiftOutput(outStr string) (models.SwiftVersionModel, error) {
	split := strings.Split(outStr, "\n")
	if len(split) == 0 {
		return models.SwiftVersionModel{}, fmt.Errorf("failed to parse swift version output (%s)", outStr)
	}

	version := split[0]
	target := split[1]

	return models.SwiftVersionModel{
		Version: version,
		Target:  target,
	}, nil
}
