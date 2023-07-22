package new

import (
	"github.com/emberfarkas/goctl/internal/migrate/clap"
	"github.com/emberfarkas/goctl/internal/migrate/utils"
	"github.com/spf13/cobra"
	"time"
)

var (
	name string
)

var Cmd = &cobra.Command{
	Use:   "new",
	Short: "通用模板新建",
	Long:  `创建migrate模板`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cmdParse clap.YamlConfig
		if clap.GenPath != "" {
			if configFileParams, err := clap.LoadConfigFile(clap.GenPath); err == nil && configFileParams != nil {
				cmdParse = *configFileParams
			}
		}
		seq := false
		seqDigits := 6
		return utils.CreateCmd(cmdParse.Path, time.Now(), utils.DefaultTimeFormat, name, "sql", seq, seqDigits, true)
	},
}

func init() {
	Cmd.Flags().StringVar(&name, "name", "default", "file name")
}
