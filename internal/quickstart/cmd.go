package quickstart

import (
	"context"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:  "quickstart",
		Long: "quickstart ",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context())
		},
	}
	name string // 名称
	out  string // 输出
)

func init() {

	// Here you will define your flags and configuration settings.
	Cmd.Flags().StringVarP(&name, "name", "n", "example", "example")
	Cmd.Flags().StringVarP(&out, "out", "o", "./Book1.xlsx", "freetoke, stt")
}

func run(ctx context.Context) error {

	return nil
}
