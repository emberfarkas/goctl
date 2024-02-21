package batchnew

import (
	"gitee.com/teacherming/keygen"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var offset int // 起始值
var limit int  // 数量

var Cmd = &cobra.Command{
	Use:   "batchnew",
	Short: "通用模板新建",
	Long:  `创建migrate模板`,
	RunE: func(cmd *cobra.Command, args []string) error {
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
		ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
		for i := offset; i < limit; i++ {
			key, err := km.GetKey(keygen.PurposeBIP44, keygen.CoinTypeETH, 1, 0, uint32(i))
			if err != nil {
				return err
			}
			priv, err := key.PrivateKey()
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
		return nil
	},
}

func init() {
	Cmd.Flags().IntVar(&offset, "offset", 0, "起始索引")
	Cmd.Flags().IntVar(&limit, "limit", 1000, "限制")
}
