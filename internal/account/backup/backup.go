package backup

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "backup",
	Short: "通用模板新建",
	Long:  `创建migrate模板`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
