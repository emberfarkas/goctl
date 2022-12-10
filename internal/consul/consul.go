package consul

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "consul",
	Short: "consul相关辅助工具",
	Long:  `一些批处理consul的工具`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// Cmd.PersistentFlags().StringVarP(&url, "url", "u", "amqp://admin:admin@127.0.0.1:5672/", "new account")
}

func run(ctx context.Context) error {
	ret, err := services(ctx)
	if err != nil {
		return err
	}
	for key, _ := range ret {
		if err = deregister(ctx, key); err != nil {
			return err
		}
	}
	return nil
}

func services(ctx context.Context) (ret map[string]*json.RawMessage, err error) {
	_, err = req.R().SetResult(&ret).Get("http://121.36.71.137:8500/v1/agent/services")
	if err != nil {
		return
	}
	return
}

func deregister(ctx context.Context, key string) (err error) {
	_, err = req.R().Put(fmt.Sprintf("http://121.36.71.137:8500/v1/agent/service/deregister/%s", key))
	if err != nil {
		return
	}
	return
}
