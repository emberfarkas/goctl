package account

import (
	"context"
	"io/fs"
	"io/ioutil"
	"math"

	"gitee.com/teacherming/keygen"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	pkgcrypto "github.com/go-bamboo/pkg/crypto"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip39"
)

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:   "account",
		Short: "账号",
		Long:  `产生账号相关的辅助工具`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context())
		},
	}
	flag   string // 标签，判断功能
	priv   string // 私钥
	offset int    // 起始值
	limit  int    // 数量
)

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	Cmd.Flags().StringVarP(&flag, "flag", "f", "", "标签: `m` 助记词")
	Cmd.Flags().StringVar(&priv, "priv", "", "私钥")
	Cmd.Flags().IntVar(&offset, "offset", 0, "起始索引")
	Cmd.Flags().IntVar(&limit, "limit", 1000, "限制")
}

func run(ctx context.Context) error {
	switch flag {
	case "k":
		_, priv, err := pkgcrypto.RSAGen()
		if err != nil {
			return err
		}
		return ioutil.WriteFile("key", []byte(priv), fs.ModePerm)
	case "m":
		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ := bip39.NewMnemonic(entropy)
		return ioutil.WriteFile("mnemonic", []byte(mnemonic), fs.ModePerm)
	case "priv":
		privKey, err := crypto.HexToECDSA(priv)
		if err != nil {
			return err
		}
		ks := keystore.NewKeyStore("keystore", int(math.Pow(2, 7)), int(math.Pow(2, 9)))
		_, err = ks.ImportECDSA(privKey, "123456")
		if err != nil {
			return err
		}
	default:
		priv, err := ioutil.ReadFile("key")
		if err != nil {
			return err
		}
		mnemonic, err := ioutil.ReadFile("mnemonic")
		if err != nil {
			return err
		}
		km, err := keygen.NewKeyManager(128, string(priv), string(mnemonic))
		if err != nil {
			return err
		}
		ks := keystore.NewKeyStore("keystore", int(math.Pow(2, 7)), int(math.Pow(2, 9)))
		for i := offset; i < limit; i++ {
			key, err := km.GetKey(keygen.PurposeBIP44, keygen.CoinTypeETH, 1, 0, uint32(i))
			if err != nil {
				return err
			}
			_, _, priv, err := key.EncodeEth()
			if err != nil {
				return err
			}
			privKey, err := crypto.HexToECDSA(priv)
			if err != nil {
				return err
			}
			_, err = ks.ImportECDSA(privKey, "123456")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
