package version

import (
	"fmt"
	"github.com/emberfarkas/goctl/internal/migrate/clap"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  `version for migrate`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("migrate[%v]", clap.Version)
		return nil
	},
}
