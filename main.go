package main

import (
	"fmt"
	"github.com/emberfarkas/goctl/internal/account"
	"github.com/emberfarkas/goctl/internal/benchmark"
	"github.com/emberfarkas/goctl/internal/binlog"
	"github.com/emberfarkas/goctl/internal/codegen"
	"github.com/emberfarkas/goctl/internal/consul"
	"github.com/emberfarkas/goctl/internal/diff"
	"github.com/emberfarkas/goctl/internal/eth"
	"github.com/emberfarkas/goctl/internal/gormgen"
	"github.com/emberfarkas/goctl/internal/leveldb"
	"github.com/emberfarkas/goctl/internal/migrate"
	"github.com/emberfarkas/goctl/internal/pdfcov"
	"github.com/emberfarkas/goctl/internal/ss"
	"github.com/emberfarkas/goctl/internal/telegram"
	"github.com/go-bamboo/pkg/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goctl",
	Short: "工具",
	Long:  `测试，辅助相关的工具`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}
var cfgFile string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//
	rootCmd.AddCommand(account.Cmd)
	rootCmd.AddCommand(benchmark.Cmd)
	rootCmd.AddCommand(consul.Cmd)
	rootCmd.AddCommand(migrate.Cmd)
	rootCmd.AddCommand(ss.Cmd)
	rootCmd.AddCommand(codegen.Cmd)
	rootCmd.AddCommand(pdfcov.Cmd)
	rootCmd.AddCommand(binlog.Cmd)
	rootCmd.AddCommand(eth.Cmd)
	rootCmd.AddCommand(telegram.Cmd)
	rootCmd.AddCommand(gormgen.Cmd)
	rootCmd.AddCommand(diff.Cmd)
	rootCmd.AddCommand(leveldb.Cmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	fmt.Print(cfgFile)
}
