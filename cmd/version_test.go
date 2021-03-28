package cmd

import (
	"testing"
)

func Test_GetCommand(t *testing.T) {
	_, err := executeCommand(Up(), "--name=test")
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}
