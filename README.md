# IQ protocol go SDK 

This repository contains the go bindings to interact with IQ protocol.

## Setting up the environment

To use this repository you'll need a Go environment set up. See the [official Go documentation](https://golang.org/doc/install) for help with that.

To run integration tests, Docker will be needed. Please refer to [Docker official documentation](https://docs.docker.com/get-docker/) for help with that.


## Running tests

To run unit tests, run:

```
go run mage.go test
```

To run integration tests, run:

```
go run mage.go testintegration
```

