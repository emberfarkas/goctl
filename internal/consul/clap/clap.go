package clap

import (
	"github.com/go-bamboo/pkg/registry"
	"gopkg.in/yaml.v3"
	"os"
)

const Version = "1.0.1"

type Config struct {
	Reg *registry.Conf `yaml:"reg"`
}

// LoadConfigFile load config file from path
func LoadConfigFile(path string) (*registry.Conf, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() // nolint
	var config Config
	if cmdErr := yaml.NewDecoder(file).Decode(&config); cmdErr != nil {
		return nil, cmdErr
	}
	return config.Reg, nil
}
