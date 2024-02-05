package clap

import (
	"github.com/go-bamboo/pkg/log"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	// DefaultOutPath default path
	DefaultOutPath = "./internal/dao/query"
)

// CmdParams is command line parameters
type CmdParams struct {
	DSN               string   `yaml:"dsn" validate:"required"`                  // consult[https://gorm.io/docs/connecting_to_the_database.html]"
	DB                string   `yaml:"db" validate:"required"`                   // input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
	Tables            []string `yaml:"tables" validate:"required,dive,required"` // enter the required data table or leave it blank
	OnlyModel         bool     `yaml:"onlyModel"`                                // only generate model
	OutPath           string   `yaml:"outPath" validate:"required"`              // specify a directory for output
	OutFile           string   `yaml:"outFile"`                                  // query code file name, default: gen.go
	WithUnitTest      bool     `yaml:"withUnitTest"`                             // generate unit test for query code
	ModelPkgName      string   `yaml:"modelPkgName"`                             // generated model code's package name
	FieldNullable     bool     `yaml:"fieldNullable"`                            // generate with pointer when field is nullable
	FieldWithIndexTag bool     `yaml:"fieldWithIndexTag"`                        // generate field with gorm index tag
	FieldWithTypeTag  bool     `yaml:"fieldWithTypeTag"`                         // generate field with gorm column type tag
	FieldSignable     bool     `yaml:"fieldSignable"`                            // detect integer field's unsigned type, adjust generated data type
}

// YamlConfig is yaml config struct
type YamlConfig struct {
	Version  string     `yaml:"version"`  //
	Database *CmdParams `yaml:"database"` //
}

type Config struct {
	GenGorm YamlConfig `yaml:"genGorm"`
}

// LoadConfigFile load config file from path
func LoadConfigFile(path string) (*CmdParams, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() // nolint
	var config Config
	if cmdErr := yaml.NewDecoder(file).Decode(&config); cmdErr != nil {
		return nil, cmdErr
	}
	log.Debugf("dsn: %v, db: %v", config.GenGorm.Database.DSN, config.GenGorm.Database.DB)
	return config.GenGorm.Database, nil
}
