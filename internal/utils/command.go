package utils

import (
	"os/exec"
	"reflect"
)

// IsNil checks a variable is nil
func IsNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

// IsCommandExist returns true if `cmd` is exist in OS
func IsCommandExist(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// ExecuteCommand executes command and return results
func ExecuteCommand(cmd string) ([]interface{}, error) {
	return nil, nil
}
