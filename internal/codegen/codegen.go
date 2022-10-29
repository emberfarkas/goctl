package codegen

import (
	"context"

	"github.com/go-bamboo/pkg/log"
	"github.com/spf13/cobra"
)

var (
	dsn         string
	androidpath string
	iospath     string
	frontpath   string
	adminpath   string
	backpath    string
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "codegen",
	Short: "初始后台模板",
	Long:  `初始后台模板，用来作为`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	Cmd.Flags().StringVar(&dsn, "dsn", "root:123456@tcp(127.0.0.1:3306)/information_schema?charset=utf8mb4&parseTime=true", "数据库dsn")
	Cmd.Flags().StringVar(&adminpath, "adminpath", "admin-ui/src", "后台路径")
	Cmd.Flags().StringVar(&backpath, "backpath", "app/service", "后端路径")
}

func run(ctx context.Context) error {
	if adminpath == "" {
		log.Fatal("adminpath is null")
	}
	if backpath == "" {
		log.Fatal("backpath is null")
	}
	w, err := makeWizard()
	if err != nil {
		return err
	}
	return w.run(ctx)
}
