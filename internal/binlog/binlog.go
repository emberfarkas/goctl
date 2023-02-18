package binlog

import (
	"context"
	"os"
	"strings"

	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:   "binlog",
		Short: "binlog相关辅助工具",
		Long:  `一些批处理binlog的工具`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context())
		},
	}
	input  string
	offset int64
	output string
)

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	Cmd.Flags().StringVarP(&input, "input", "i", "", "path for file")
	Cmd.Flags().Int64VarP(&offset, "offset", "n", 0, "offset for file")
	Cmd.Flags().StringVarP(&output, "output", "o", "", "path for file")
}

func run(ctx context.Context) error {
	var sb strings.Builder
	p := replication.NewBinlogParser()
	f := func(e *replication.BinlogEvent) error {
		e.Dump(&sb)
		return nil
	}
	if err := p.ParseFile(input, offset, f); err != nil {
		println(err.Error())
		return err
	}
	fn, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0)
	if err != nil {
		return err
	}
	defer fn.Close()
	c := sb.String()
	_, err = fn.WriteString(c)
	if err != nil {
		return err
	}
	return nil
}
