package version

import (
	"fmt"
	"github.com/emberfarkas/goctl/internal/diff/clap"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "version",
		Short: "version相关",
		Long:  `比较version文件不同详情`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("diff[%v]", clap.Version)
			return nil
		},
	}
)
