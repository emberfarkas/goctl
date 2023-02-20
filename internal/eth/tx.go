package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-bamboo/pkg/log"
	"github.com/spf13/cobra"
)

// 这个工具主要用来测试eth相关的借口

// Cmd represents the config command
var (
	txCmd = &cobra.Command{
		Use:   "tx",
		Short: "tx相关",
		Long:  `获取交易详情`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return getTx(cmd.Context())
		},
	}
	hash string // 交易hash
)

func init() {

	// Here you will define your flags and configuration settings.

	//txCmd.Flags().StringVarP(&hash, "hash", "h", "", "标签: `m` 助记词")
}

func getTx(ctx context.Context) error {
	hash = "0xb5519dc9333aaed59898de7093946dc22c69f240a40c5625e67c02b12749c85c"
	log.Infof("hash: %v", hash)
	rpc, err := ethclient.Dial("https://bsc-dataseed4.ninicoin.io")
	if err != nil {
		return err
	}
	tx, _, err := rpc.TransactionByHash(ctx, common.HexToHash(hash))
	if err != nil {
		return err
	}
	log.Debugf("from: %v", tx.To())
	log.Debugf("to: %v", tx.To)

	return nil
}
