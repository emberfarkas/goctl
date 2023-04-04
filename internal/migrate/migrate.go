package migrate

import (
	"errors"
	"os"
	"time"

	"github.com/go-bamboo/pkg/log"
	commonds "github.com/go-bamboo/pkg/migrate"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	// _ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	errInvalidSequenceWidth     = errors.New("Digits must be positive")
	errIncompatibleSeqAndFormat = errors.New("The seq and format options are mutually exclusive")
	errInvalidTimeFormat        = errors.New("Time format may not be empty")
	errNotSupport               = errors.New("not support cmd")
)

// DefaultOutPath default path
const (
	defaultPath       = "file://./migrations"
	defaultTimeFormat = "20060102150405"
)

var (
	name string

	genPath     string
	sourceURL   string
	path        string
	databaseURL string
)

// YamlConfig is yaml config struct
type YamlConfig struct {
	Version     string `yaml:"version"` //
	SourceURL   string `yaml:"source"`
	Path        string `yaml:"path"` // specify a directory for migrates
	DatabaseURL string `yaml:"database"`
}

type Config struct {
	Migrate YamlConfig `yaml:"migrate"`
}

// loadConfigFile load config file from path
func loadConfigFile(path string) (*YamlConfig, error) {
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

// Cmd represents the new command
var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate模块",
	Long:  `migrate相关的统计`,
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "通用模板新建",
	Long:  `创建migrate模板`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cmdParse YamlConfig
		if genPath != "" {
			if configFileParams, err := loadConfigFile(genPath); err == nil && configFileParams != nil {
				cmdParse = *configFileParams
			}
		}
		seq := false
		seqDigits := 6
		return commonds.CreateCmd(cmdParse.Path, time.Now(), defaultTimeFormat, name, "sql", seq, seqDigits, true)
	},
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "通用模板新建",
	Long:  `创建migrate模板`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cmdParse YamlConfig
		if genPath != "" {
			if configFileParams, err := loadConfigFile(genPath); err == nil && configFileParams != nil {
				cmdParse = *configFileParams
			}
		}
		seq := false
		seqDigits := 6
		return commonds.CreateCmd(cmdParse.Path, time.Now(), defaultTimeFormat, name, "sql", seq, seqDigits, true)
	},
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "更新数据库",
	Long:  `更新数据库`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cmdParse YamlConfig
		if genPath != "" {
			if configFileParams, err := loadConfigFile(genPath); err == nil && configFileParams != nil {
				cmdParse = *configFileParams
			}
		}
		// cmd first
		if sourceURL != "" {
			cmdParse.SourceURL = sourceURL
		}
		if sourceURL == "" && path != "" {
			cmdParse.SourceURL = path
		}
		if cmdParse.SourceURL == "" {
			cmdParse.SourceURL = defaultPath
		}
		if databaseURL != "" {
			cmdParse.DatabaseURL = databaseURL
		}
		m, err := migrate.New(cmdParse.SourceURL, cmdParse.DatabaseURL)
		defer func() {
			if err == nil {
				if _, err := m.Close(); err != nil {
					log.Error(err)
				}
			}
		}()
		if err != nil {
			return err
		}
		if err = m.Up(); err != nil {
			if err.Error() == "no change" {
			}
			return err
		}
		return nil
	},
}

func init() {
	Cmd.AddCommand(newCmd, exportCmd, upCmd)
	Cmd.PersistentFlags().StringVar(&genPath, "c", "", "is path for gen.yml")

	// Here you will define your flags and configuration settings.
	newCmd.Flags().StringVar(&name, "name", "default", "file name")

	// exportCmd.Flags().StringVar(&dsn, "dsn", "", "数据库链接")
	// exportCmd.Flags().StringVar(&tableList, "tables", "", "表")

	upCmd.Flags().StringVar(&sourceURL, "source", "", "consult[https://gorm.io/docs/connecting_to_the_database.html]")
	upCmd.Flags().StringVar(&path, "path", "", "specify a directory for output")
	upCmd.Flags().StringVar(&databaseURL, "db", "", "input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]")
}
