package priv

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var (
	path string
	priv string
	pass string
)

var Cmd = &cobra.Command{
	Use:   "priv",
	Short: "通用模板新建",
	Long:  `根据私钥创建`,
	RunE: func(cmd *cobra.Command, args []string) error {
		privKey, err := crypto.HexToECDSA(priv)
		if err != nil {
			return err
		}
		ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
		_, err = ks.ImportECDSA(privKey, pass)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&path, "path", "default", "file name")
	Cmd.Flags().StringVar(&priv, "priv", "default", "file name")
	Cmd.Flags().StringVar(&pass, "pass", "123456", "file name")
}
