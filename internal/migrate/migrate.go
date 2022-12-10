package migrate

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/migrate"
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
const DefaultOutPath = "./migrations"

var (
	defaultTimeFormat = "20060102150405"

	name string

	genPath   string
	dsn       string
	db        string
	tableList string
	outPath   string
)

// CmdParams is command line parameters
type CmdParams struct {
	DSN     string `yaml:"dsn"`     // consult[https://gorm.io/docs/connecting_to_the_database.html]"
	DB      string `yaml:"db"`      // input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
	OutPath string `yaml:"outPath"` // specify a directory for output
}

// YamlConfig is yaml config struct
type YamlConfig struct {
	Version  string     `yaml:"version"`  //
	Database *CmdParams `yaml:"database"` //
}

type Config struct {
	GenGorm YamlConfig `yaml:"migrate"`
}

// loadConfigFile load config file from path
func loadConfigFile(path string) (*CmdParams, error) {
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
		seq := false
		seqDigits := 6
		return createCmd("./", time.Now(), defaultTimeFormat, name, "sql", seq, seqDigits, true)
	},
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "通用模板新建",
	Long:  `创建migrate模板`,
	RunE: func(cmd *cobra.Command, args []string) error {
		seq := false
		seqDigits := 6
		return createCmd("./", time.Now(), defaultTimeFormat, name, "sql", seq, seqDigits, true)
	},
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "更新数据库",
	Long:  `更新数据库`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cmdParse CmdParams
		if genPath != "" {
			if configFileParams, err := loadConfigFile(genPath); err == nil && configFileParams != nil {
				cmdParse = *configFileParams
			}
		}
		// cmd first
		if dsn != "" {
			cmdParse.DSN = dsn
		}
		if db != "" {
			cmdParse.DB = db
		}
		if outPath != "" {
			cmdParse.OutPath = outPath
		}
		//
		sourceURL := fmt.Sprintf("file://%v", cmdParse.OutPath)
		mf := &migrate.Config{
			Driver:    cmdParse.DB,
			Source:    cmdParse.DSN,
			SourceURL: sourceURL,
		}
		if err := migrate.MigrateUp(mf); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	Cmd.AddCommand(newCmd, exportCmd, upCmd)

	// Here you will define your flags and configuration settings.

	newCmd.Flags().StringVar(&name, "name", "default", "file name")

	exportCmd.Flags().StringVar(&dsn, "dsn", "", "数据库链接")
	exportCmd.Flags().StringVar(&tableList, "tables", "", "表")

	upCmd.Flags().StringVar(&genPath, "c", "", "is path for gen.yml")
	upCmd.Flags().StringVar(&dsn, "dsn", "", "consult[https://gorm.io/docs/connecting_to_the_database.html]")
	upCmd.Flags().StringVar(&db, "db", "", "input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]")
	upCmd.Flags().StringVar(&outPath, "outPath", DefaultOutPath, "specify a directory for output")
}
