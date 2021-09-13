// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Runs the test suite against the repo
func Test() error {
	return sh.RunV("go", "test", "--short", "-race", "-cover", "./...")
}
