package new

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/cobra"
	"log"
)

var (
	path string
	pass string
)

var Cmd = &cobra.Command{
	Use:   "new",
	Short: "通用模板新建",
	Long:  `创建migrate模板`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
		account, err := ks.NewAccount(pass)
		if err != nil {
			return err
		}
		log.Printf("%v", account.Address.Hex())
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&path, "path", "default", "file name")
	Cmd.Flags().StringVar(&pass, "pass", "123456", "password")
}
