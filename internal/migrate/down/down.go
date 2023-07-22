package down

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "down",
	Short: "通用模板新建",
	Long:  `创建migrate模板`,
}
