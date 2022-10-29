package eth

import (
	"context"
	"github.com/go-bamboo/pkg/log"
	"github.com/onrik/ethrpc"
	"github.com/spf13/cobra"
)

// 这个工具主要用来测试eth相关的借口

// Cmd represents the config command
var (
	codeCmd = &cobra.Command{
		Use:   "code",
		Short: "code相关",
		Long:  `获取交易详情`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return getCode(cmd.Context())
		},
	}
	//hash string // 交易hash
)

func init() {

	// Here you will define your flags and configuration settings.

	//txCmd.Flags().StringVarP(&hash, "hash", "h", "", "标签: `m` 助记词")
}

func getCode(ctx context.Context) error {
	rpc := ethrpc.New("https://bsc-dataseed4.ninicoin.io")
	code, err := rpc.EthGetCode("0x309a174c27da0f03e6c36617f57d9560c5895991", "latest")
	if err != nil {
		return err
	}
	log.Debugf("code: %v", code)
	return nil
}
