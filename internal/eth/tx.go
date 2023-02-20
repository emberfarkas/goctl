package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
			return transfer(cmd.Context())
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

func transfer(ctx context.Context) (err error) {
	client, err := ethclient.Dial("https://endpoints.omniatech.io/v1/bsc/testnet/public")
	if err != nil {
		return err
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return
	}
	log.Infof("chainID[%v]", chainID)
	privKey, err := crypto.HexToECDSA("95884f665f4cf15b77a75017b64f9ff7df93b565fb9f3415b4cb352bd627141e")
	if err != nil {
		return err
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("invalid key")
	}
	from := crypto.PubkeyToAddress(*pubKeyECDSA)
	no, err := client.NonceAt(ctx, from, nil)
	if err != nil {
		return
	}
	log.Infof("nonce: %v\n", no)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return
	}
	tokenAddress := common.HexToAddress("0x62AB07f83cc62f4bf940D0330f6019588Deed13e")
	value := big.NewInt(int64(math.Pow10(18) * 0.001))
	data := []byte("")
	tx := types.NewTransaction(uint64(no), tokenAddress, value, 21000, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return
	}
	log.Debugf("hash: %v", signedTx.Hash())
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return
	}
	return
}
