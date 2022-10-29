package consul

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-bamboo/pkg/net/http"
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
	agent, err := http.NewSuperAgent(http.Timeout(5 * time.Second))
	if err != nil {
		return
	}
	_, _, errs := agent.Get("http://121.36.71.137:8500/v1/agent/services").EndStruct(&ret)
	if len(errs) > 0 {
		err = errs[0]
		return
	}
	return
}

func deregister(ctx context.Context, key string) (err error) {
	agent, err := http.NewSuperAgent(http.Timeout(5 * time.Second))
	if err != nil {
		return
	}
	_, _, errs := agent.Put(fmt.Sprintf("http://121.36.71.137:8500/v1/agent/service/deregister/%s", key)).End()
	if len(errs) > 0 {
		err = errs[0]
		return
	}
	return
}
