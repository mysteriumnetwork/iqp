// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Runs the unit test suite against the repo
func Test() error {
	return sh.RunV("go", "test", "--short", "-race", "-cover", "./...")
}

func TestIntegration() error {
	err := TestEIP155Client()
	if err != nil {
		return err
	}

	err = TestPostgresStorage()
	return err
}

var postgresComposeFile = "./account/docker-compose.integration.yml"

// Runs the integration test for postgres storage
func TestPostgresStorage() error {
	defer stopCompose(postgresComposeFile)
	_, err := startCompose(postgresComposeFile)
	if err != nil {
		return err
	}

	return sh.RunV("go", "test", "-race", "-cover", "-run", "^TestPostgresStorage_integration$", "./account")
}

// Runs the integration test for eip155 blockchain client
func TestEIP155Client() error {
	defer stopCompose(eip155TestComposeFile)
	_, err := startCompose(eip155TestComposeFile)
	if err != nil {
		return err
	}

	return sh.RunV("go", "test", "-race", "-cover", "./blockchain/eip155")
}

var eip155TestComposeFile = "./blockchain/eip155/docker-compose.integration.yml"

func startCompose(filepath string) (string, error) {
	return sh.Output("docker-compose", "-f", filepath, "up", "-d")
}

func stopCompose(filepath string) (string, error) {
	return sh.Output("docker-compose", "-f", filepath, "down")
}
