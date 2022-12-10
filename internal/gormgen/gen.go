package gormgen

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/store/gormx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"gorm.io/gen"
)

const (
	// DefaultOutPath default path
	DefaultOutPath = "./internal/dao/query"
)

// CmdParams is command line parameters
type CmdParams struct {
	DSN               string   `yaml:"dsn"`               // consult[https://gorm.io/docs/connecting_to_the_database.html]"
	DB                string   `yaml:"db"`                // input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
	Tables            []string `yaml:"tables"`            // enter the required data table or leave it blank
	OnlyModel         bool     `yaml:"onlyModel"`         // only generate model
	OutPath           string   `yaml:"outPath"`           // specify a directory for output
	OutFile           string   `yaml:"outFile"`           // query code file name, default: gen.go
	WithUnitTest      bool     `yaml:"withUnitTest"`      // generate unit test for query code
	ModelPkgName      string   `yaml:"modelPkgName"`      // generated model code's package name
	FieldNullable     bool     `yaml:"fieldNullable"`     // generate with pointer when field is nullable
	FieldWithIndexTag bool     `yaml:"fieldWithIndexTag"` // generate field with gorm index tag
	FieldWithTypeTag  bool     `yaml:"fieldWithTypeTag"`  // generate field with gorm column type tag
	FieldSignable     bool     `yaml:"fieldSignable"`     // detect integer field's unsigned type, adjust generated data type
}

// YamlConfig is yaml config struct
type YamlConfig struct {
	Version  string     `yaml:"version"`  //
	Database *CmdParams `yaml:"database"` //
}

type Config struct {
	GenGorm YamlConfig `yaml:"genGorm"`
}

// genModels is gorm/gen generated models
func genModels(g *gen.Generator, db *gormx.DB, tables []string) (models []interface{}, err error) {
	var tablesList []string
	if len(tables) == 0 {
		// Execute tasks for all tables in the database
		tablesList, err = db.Migrator().GetTables()
		if err != nil {
			return nil, fmt.Errorf("GORM migrator get all tables fail: %w", err)
		}
	} else {
		tablesList = tables
	}

	// Execute some data table tasks
	models = make([]interface{}, len(tablesList))
	for i, tableName := range tablesList {
		models[i] = g.GenerateModel(tableName)
	}
	return models, nil
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

// argParse is parser for cmd
func argParse() *CmdParams {
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
	if tableList != "" {
		cmdParse.Tables = strings.Split(tableList, ",")
	}
	if onlyModel {
		cmdParse.OnlyModel = true
	}
	if outPath != DefaultOutPath {
		cmdParse.OutPath = outPath
	}
	if outFile != "" {
		cmdParse.OutFile = outFile
	}
	if withUnitTest {
		cmdParse.WithUnitTest = withUnitTest
	}
	if modelPkgName != "" {
		cmdParse.ModelPkgName = modelPkgName
	}
	if fieldNullable {
		cmdParse.FieldNullable = fieldNullable
	}
	if fieldWithIndexTag {
		cmdParse.FieldWithIndexTag = fieldWithIndexTag
	}
	if fieldWithTypeTag {
		cmdParse.FieldWithTypeTag = fieldWithTypeTag
	}
	if fieldSignable {
		cmdParse.FieldSignable = fieldSignable
	}
	return &cmdParse
}

var (
	genPath           string
	dsn               string
	db                string
	tableList         string
	onlyModel         bool
	outPath           string
	outFile           string
	withUnitTest      bool
	modelPkgName      string
	fieldNullable     bool
	fieldWithIndexTag bool
	fieldWithTypeTag  bool
	fieldSignable     bool
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "gormgen",
	Short: "初始Gorm模板",
	Long:  `初始Gorm模板，用来作为`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// choose is file or flag
	Cmd.Flags().StringVar(&genPath, "c", "", "is path for gen.yml")
	Cmd.Flags().StringVar(&dsn, "dsn", "", "consult[https://gorm.io/docs/connecting_to_the_database.html]")
	Cmd.Flags().StringVar(&db, "db", "", "input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]")
	Cmd.Flags().StringVar(&tableList, "tables", "", "enter the required data table or leave it blank")
	Cmd.Flags().BoolVar(&onlyModel, "onlyModel", false, "only generate models (without query file)")
	Cmd.Flags().StringVar(&outPath, "outPath", DefaultOutPath, "specify a directory for output")
	Cmd.Flags().StringVar(&outFile, "outFile", "", "query code file name, default: gen.go")
	Cmd.Flags().BoolVar(&withUnitTest, "withUnitTest", false, "generate unit test for query code")
	Cmd.Flags().StringVar(&modelPkgName, "modelPkgName", "", "generated model code's package name")
	Cmd.Flags().BoolVar(&fieldNullable, "fieldNullable", false, "generate with pointer when field is nullable")
	Cmd.Flags().BoolVar(&fieldWithIndexTag, "fieldWithIndexTag", false, "generate field with gorm index tag")
	Cmd.Flags().BoolVar(&fieldWithTypeTag, "fieldWithTypeTag", false, "generate field with gorm column type tag")
	Cmd.Flags().BoolVar(&fieldSignable, "fieldSignable", false, "detect integer field's unsigned type, adjust generated data type")
}

func run(ctx context.Context) error {
	// cmdParse
	config := argParse()
	if config == nil {
		log.Fatal("parse config fail")
	}

	db := gormx.MustNew(&gormx.Conf{
		Driver:   config.DB,
		Source:   config.DSN,
		LogLevel: 4,
	})

	g := gen.NewGenerator(gen.Config{
		OutPath:           config.OutPath,
		OutFile:           config.OutFile,
		ModelPkgPath:      config.ModelPkgName,
		WithUnitTest:      config.WithUnitTest,
		FieldNullable:     config.FieldNullable,
		FieldWithIndexTag: config.FieldWithIndexTag,
		FieldWithTypeTag:  config.FieldWithTypeTag,
		FieldSignable:     config.FieldSignable,
	})

	g.UseDB(db)

	models, err := genModels(g, db, config.Tables)
	if err != nil {
		log.Fatal("get tables info fail:", err)
	}

	if !config.OnlyModel {
		g.ApplyBasic(models...)
	}

	g.Execute()
	return nil
}
