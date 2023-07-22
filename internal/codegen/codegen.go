package codegen

import (
	"github.com/emberfarkas/goctl/internal/codegen/backend/dao"
	"github.com/emberfarkas/goctl/internal/codegen/backend/model"
	"github.com/emberfarkas/goctl/internal/codegen/utils"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "codegen",
	Short: "初始后台模板",
	Long:  `初始后台模板，用来作为`,
}

func init() {

	// Here you will define your flags and configuration settings.
	Cmd.PersistentFlags().StringVar(&utils.DSN, "dsn", "root:123456@tcp(127.0.0.1:3306)/information_schema?charset=utf8mb4&parseTime=true", "数据库dsn")

	Cmd.AddCommand(dao.Cmd)
	Cmd.AddCommand(model.Cmd)
}
