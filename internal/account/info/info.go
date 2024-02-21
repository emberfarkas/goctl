package info

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"log"
)

var (
	path string
	pass string
)

var Cmd = &cobra.Command{
	Use:   "info",
	Short: "信息",
	Long:  `查看所有信息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
		for _, account := range ks.Accounts() {
			keyJSON, err := ks.Export(account, pass, pass)
			if err != nil {
				return err
			}
			key, err := keystore.DecryptKey(keyJSON, pass)
			if err != nil {
				return err
			}
			pkey := hexutil.Encode(key.PrivateKey.D.Bytes())
			log.Printf("pkey: %v", pkey)
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&path, "path", "default", "file name")
	Cmd.Flags().StringVar(&pass, "pass", "123456", "password")
}
