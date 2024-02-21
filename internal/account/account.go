package account

import (
	"github.com/emberfarkas/goctl/internal/account/backup"
	"github.com/emberfarkas/goctl/internal/account/batchnew"
	"github.com/emberfarkas/goctl/internal/account/info"
	"github.com/emberfarkas/goctl/internal/account/key"
	"github.com/emberfarkas/goctl/internal/account/mnemonic"
	"github.com/emberfarkas/goctl/internal/account/new"
	"github.com/emberfarkas/goctl/internal/account/priv"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:   "account",
		Short: "账号",
		Long:  `产生账号相关的辅助工具`,
	}
)

func init() {

	// Here you will define your flags and configuration settings.

	Cmd.AddCommand(key.Cmd)
	Cmd.AddCommand(batchnew.Cmd)
	Cmd.AddCommand(backup.Cmd)
	Cmd.AddCommand(mnemonic.Cmd)
	Cmd.AddCommand(new.Cmd)
	Cmd.AddCommand(priv.Cmd)
	Cmd.AddCommand(info.Cmd)
}
