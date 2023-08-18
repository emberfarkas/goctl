package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

// go build -ldflags "-X version.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// Branch is current branch name the code is built off.
	Branch string
	// Revision is the short commit hash of source tree.
	Revision string
	// BuildDate is the date when the binary was built.
	BuildDate string
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "version",
	Short: "version相关辅助工具",
	Long:  `打印版本信息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("version: %v-%v[%v]\n", Version, Revision, BuildDate)
		return nil
	},
}
