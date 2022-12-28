package diff

import "github.com/spf13/cobra"

// 这个工具主要用来对比wav文件

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:   "diff",
		Short: "diff相关",
		Long:  `diff相关的辅助工具`,
	}
)

func init() {
	Cmd.AddCommand(wavCmd)
}
