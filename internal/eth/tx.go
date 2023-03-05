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
	"github.com/tidwall/sjson"
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
	//bsc := "https://bsc-dataseed4.ninicoin.io"
	eth := "https://ethereum.blockpi.network/v1/rpc/public"
	hash = "0x64dff047841b2138a9f5200ce319d50fb37aec1ef8489ef34fa577de461d642d"
	rpc, err := ethclient.Dial(eth)
	if err != nil {
		return err
	}
	tx, _, err := rpc.TransactionByHash(ctx, common.HexToHash(hash))
	if err != nil {
		return err
	}
	log.Debugf("%v", tx.To())
	v, r, s := tx.RawSignatureValues()
	log.Debugf("v[%v], r[%v], s[%v]", v, r, s)
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
	privKey, err := crypto.HexToECDSA("c0247b7f40e5c29a405eafbd8316c7d7ff904fbf04fc95704debcdd9214bc8e2")
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
	tx := types.NewTransaction(no, tokenAddress, value, 21000, gasPrice, data)
	signer := types.NewEIP155Signer(chainID)
	signedTx97, err := types.SignTx(tx, signer, privKey)
	if err != nil {
		return
	}
	signedTx97.MarshalJSON()
	v, r, s := signedTx97.RawSignatureValues()
	log.Debugf("hash: %v, v[%v], r[%v], s[%v]", signedTx97.Hash(), v, r, s)

	txBytes, err := signedTx97.MarshalJSON()
	if err != nil {
		return err
	}
	txBytes, err = sjson.SetBytes(txBytes, "v", "0x1c")
	if err != nil {
		return err
	}
	var signedTx97Ex types.Transaction
	if err := signedTx97Ex.UnmarshalJSON(txBytes); err != nil {
		return err
	}
	err = client.SendTransaction(ctx, &signedTx97Ex)
	if err != nil {
		return
	}
	//tx, _, err := client.TransactionByHash(ctx, common.HexToHash("0x4787f538e43fa0b08849dd958a101dcb5b91200c9169cced33b6ead16d5fc507"))
	//if err != nil {
	//	return err
	//}
	//var rs rlp.Stream
	//if err := tx.DecodeRLP(&rs); err != nil {
	//	return err
	//}

	return
}

func replay(ctx context.Context) (err error) {
	client, err := ethclient.Dial("https://exchaintestrpc.okex.org")
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

	return nil
}
