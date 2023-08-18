package clap

import (
	"github.com/go-bamboo/pkg/log"
	"gopkg.in/yaml.v3"
	"os"
)

const Version = "1.0.1"

var GenPath string

// YamlConfig is yaml config struct
type YamlConfig struct {
	Version     int    `yaml:"version"` //
	SourceURL   string `yaml:"source"`
	DatabaseURL string `yaml:"database"`
}

type Config struct {
	Migrate YamlConfig `yaml:"migrate"`
}

// LoadConfigFile load config file from path
func LoadConfigFile(path string) (*YamlConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() // nolint
	var config Config
	if cmdErr := yaml.NewDecoder(file).Decode(&config); cmdErr != nil {
		return nil, cmdErr
	}
	log.Debugf("source: %v, db: %v", config.Migrate.SourceURL, config.Migrate.DatabaseURL)
	return &config.Migrate, nil
}
