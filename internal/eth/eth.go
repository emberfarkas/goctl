package eth

import "github.com/spf13/cobra"

// 这个工具主要用来测试eth相关的借口

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:   "eth",
		Short: "eth相关",
		Long:  `以太坊相关的辅助工具`,
	}
)

func init() {
	Cmd.AddCommand(txCmd)
	Cmd.AddCommand(codeCmd)
}
