package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	errors "golang.org/x/xerrors"
)

type Config struct {
	DB_IP               string
	DB_PORT             string
	DB_USER             string
	DB_PASSWORD         string
	DB_NAME             string
	BRIDGE_ADDR         string
	BootStrapIPs        []string
	RinkebyBootStrapIPs []string
	HecoBootStrapIPs    []string
	BSCBootStrapIPs     []string
	ChainNodePool       []string
}

// LoadConfig loads configuration file from path.
func LoadConfig(path string) (c *Config, err error) {
	var jsonFile *os.File
	var byteValue []byte

	fmt.Println("Path ", path)
	// Open our jsonFile
	jsonFile, err = os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		errors.Errorf(": %w", err)
		return
	}
	fmt.Println("Successfully Opened json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		errors.Errorf(": %w", err)
		return
	}
	c = &Config{}
	err = json.Unmarshal(byteValue, c)
	if err != nil {
		errors.Errorf(": %w", err)
		return
	}
	return
}
