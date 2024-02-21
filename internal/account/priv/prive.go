package priv

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var (
	priv string
)

var Cmd = &cobra.Command{
	Use:   "priv",
	Short: "通用模板新建",
	Long:  `创建migrate模板`,
	RunE: func(cmd *cobra.Command, args []string) error {
		privKey, err := crypto.HexToECDSA(priv)
		if err != nil {
			return err
		}
		ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
		_, err = ks.ImportECDSA(privKey, "123456")
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&priv, "priv", "default", "file name")
}
