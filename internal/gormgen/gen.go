package gormgen

import (
	"context"
	"fmt"
	"strings"

	"github.com/emberfarkas/goctl/internal/gormgen/clap"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/store/gormx"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"gorm.io/gen"
)

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

// argParse is parser for cmd
func argParse() *clap.CmdParams {
	var cmdParse clap.CmdParams = clap.CmdParams{
		DSN:               dsn,
		DB:                db,
		Tables:            strings.Split(tableList, ","),
		OnlyModel:         true,
		OutPath:           clap.DefaultOutPath,
		OutFile:           outFile,
		WithUnitTest:      withUnitTest,
		ModelPkgName:      modelPkgName,
		FieldNullable:     fieldNullable,
		FieldWithIndexTag: fieldWithIndexTag,
		FieldWithTypeTag:  fieldWithTypeTag,
		FieldSignable:     fieldSignable,
	}
	if genPath != "" {
		if configFileParams, err := clap.LoadConfigFile(genPath); err == nil && configFileParams != nil {
			cmdParse = *configFileParams
			return &cmdParse
		}
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
	Cmd.Flags().StringVar(&outPath, "outPath", clap.DefaultOutPath, "specify a directory for output")
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
	var validate *validator.Validate = validator.New()
	if err := validate.Struct(config); err != nil {
		return err
	}

	db := gormx.MustNew(&gormx.Conf{
		Driver: gormx.DBType(gormx.DBType_value[config.DB]),
		Source: config.DSN,
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
