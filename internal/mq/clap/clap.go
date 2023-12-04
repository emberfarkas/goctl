package clap

import (
	"github.com/go-bamboo/pkg/client/rabbitmq"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

const Version = "1.0.1"

// YamlConfig is yaml config struct
type YamlConfig struct {
	Version string               `yaml:"version"` //
	Conn    *rabbitmq.RabbitConf `yaml:"conn"`
	Names   []string             `yaml:"names"`
	Ttl     time.Time            `yaml:"ttl"`
}

type Config struct {
	RabbitmqAdmin YamlConfig `yaml:"rabbitmqAdmin"`
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
	return &config.RabbitmqAdmin, nil
}
