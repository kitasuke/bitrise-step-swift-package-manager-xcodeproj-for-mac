package utility

import (
	"testing"

	"github.com/kitasuke/go-swift/models"
)

const testSwiftVersionOutput = `Apple Swift version 4.2 (swiftlang-1000.11.37.1 clang-1000.11.45.1)
Target: x86_64-apple-darwin18.2.0`

func TestGetSwiftVersion(t *testing.T) {
	t.Log("GetSwiftVersion")
	{
		expectedVersion := models.SwiftVersionModel{
			Version: "Apple Swift version 4.2 (swiftlang-1000.11.37.1 clang-1000.11.45.1)",
			Target:  "Target: x86_64-apple-darwin18.2.0",
		}

		version, err := getSwiftVersionFromSwiftOutput(testSwiftVersionOutput)
		if err != nil {
			t.Fatalf("getSwiftVersionFromSwiftOutput got error: %s", err)
		}
		if version.Version != expectedVersion.Version {
			t.Errorf("expectedVersion.Version is wrong. want=%s, got=%s", expectedVersion.Version, version.Version)
		}
		if version.Target != expectedVersion.Target {
			t.Errorf("expectedVersion.Target is wrong. want=%s, got=%s", expectedVersion.Target, version.Target)
		}
	}
}
