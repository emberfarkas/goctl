package migrate

import (
	"errors"
	"fmt"
	"time"

	"github.com/emberfarkas/pkg/migrate"
	"github.com/spf13/cobra"
)

var (
	errInvalidSequenceWidth     = errors.New("Digits must be positive")
	errIncompatibleSeqAndFormat = errors.New("The seq and format options are mutually exclusive")
	errInvalidTimeFormat        = errors.New("Time format may not be empty")
	errNotSupport               = errors.New("not support cmd")
)

var (
	defaultTimeFormat = "20060102150405"
	name              string
	dsn               string
	user              string = "root"
	pass              string = "123456"
	host              string = "127.0.0.1"
	port              uint32 = 3306
	db                string
	table             string
	path              string
)

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
		//
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=true", user, pass, host, port, db)
		sourceURL := fmt.Sprintf("file://%v", path)
		mf := &migrate.Config{
			Driver:    "mysql",
			Source:    dsn,
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

	newCmd.Flags().StringVarP(&name, "name", "n", "default", "file name")

	exportCmd.Flags().StringVar(&dsn, "dsn", "", "数据库链接")
	exportCmd.Flags().StringVar(&table, "table", "", "表")

	upCmd.Flags().StringVarP(&user, "user", "u", "root", "用户")
	upCmd.Flags().StringVarP(&pass, "pass", "p", "123456", "密码")
	upCmd.Flags().StringVar(&host, "host", "127.0.0.1", "主机")
	upCmd.Flags().Uint32Var(&port, "port", 3306, "端口")
	upCmd.Flags().StringVar(&db, "db", "", "数据库")
	upCmd.Flags().StringVar(&path, "path", "./migrations", "路径")
}
