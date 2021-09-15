// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Runs the unit test suite against the repo
func Test() error {
	return sh.RunV("go", "test", "--short", "-race", "-cover", "./...")
}

// Runs the integration test for eip155 blockchain client
func TestEIP155Client() error {
	defer stopEIP155Compose()
	_, err := startEIP155Compose()
	if err != nil {
		return err
	}

	return sh.RunV("go", "test", "-race", "-cover", "./blockchain/eip155")
}

var eip155TestComposeFile = "./blockchain/eip155/docker-compose.integration.yml"

func startEIP155Compose() (string, error) {
	return sh.Output("docker-compose", "-f", eip155TestComposeFile, "up", "-d")
}

func stopEIP155Compose() (string, error) {
	return sh.Output("docker-compose", "-f", eip155TestComposeFile, "down")
}
