package diff

import (
	"github.com/emberfarkas/goctl/internal/diff/clap"
	"github.com/emberfarkas/goctl/internal/diff/wav"
	"github.com/spf13/cobra"
)

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
	Cmd.AddCommand(wav.Cmd)
	// Here you will define your flags and configuration settings.
	Cmd.PersistentFlags().StringVar(&clap.Src, "src", "", "对比者")
	Cmd.PersistentFlags().StringVar(&clap.Src, "dst", "", "被对比者")
}
