package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var pkgName = flag.String("pkg", "bindings", "the package name to use in generated bindings")
var contractDir = flag.String("contract-dir", "bindings/contracts", "the directory containing the contracts")
var out = flag.String("out", "bindings", "the ouput directory")

func main() {
	flag.Parse()
	smartContracts, err := loadFromLocalDir(*contractDir)
	if err != nil {
		fmt.Println("Error parsing truffle output: ", err.Error())
		os.Exit(-1)
	}

	for _, smartContract := range smartContracts {
		genCode, err := bindSmartContract(smartContract, *pkgName)
		if err != nil {
			fmt.Println("Error binding smart contract: ", err.Error())
			os.Exit(-1)
		}
		if err := writeToOutput(smartContract.ContractName, genCode, *out); err != nil {
			fmt.Println("Error writing generated code: ", err.Error())
			os.Exit(-1)
		}
	}

	fmt.Println("Bindings generated. Please remove the duplicate structs and commit the new bindings.")
}

func writeToOutput(fileName string, genCode, output string) error {
	var writer io.Writer
	if err := os.MkdirAll(output, 0755); err != nil {
		return err
	}
	file, err := os.Create(filepath.Join(output, fileName+".go"))
	if err != nil {
		return err
	}
	defer file.Close()
	writer = file
	if _, err := io.WriteString(writer, genCode); err != nil {
		return err
	}
	return nil
}

func bindSmartContract(smartContract truffleArtifact, pkgName string) (string, error) {
	genCode, err := bind.Bind([]string{smartContract.ContractName}, []string{smartContract.AbiString()}, []string{smartContract.Bytecode}, nil, pkgName, bind.LangGo, nil, nil)
	if err != nil {
		return "", err
	}
	return genCode, nil
}

func loadFromLocalDir(localDir string) ([]truffleArtifact, error) {
	files, err := ioutil.ReadDir(localDir)
	if err != nil {
		return nil, err
	}

	var artifacts []truffleArtifact
	for _, f := range files {
		reader, err := os.Open(localDir + "/" + f.Name())
		if err != nil {
			return nil, err
		}
		artifact, err := parseTruffleArtifact(reader)
		if err != nil {
			return nil, err
		}
		artifacts = append(artifacts, artifact)
	}

	return artifacts, nil
}

func parseTruffleArtifact(inputReader io.Reader) (truffleArtifact, error) {
	var output truffleArtifact
	if err := json.NewDecoder(inputReader).Decode(&output); err != nil {
		return truffleArtifact{}, err
	}
	return output, nil
}

type truffleArtifact struct {
	Bytecode     string          `json:"bytecode"`
	AbiBytes     json.RawMessage `json:"abi"`
	ContractName string          `json:"contractName"`
}

func (ta truffleArtifact) AbiString() string {
	return string(ta.AbiBytes)
}
