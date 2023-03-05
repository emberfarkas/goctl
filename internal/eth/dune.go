package eth

import (
	"context"
	"github.com/spf13/cobra"
)

// 这个工具主要用来测试eth相关的借口

// Cmd represents the config command
var (
	duneCmd = &cobra.Command{
		Use:   "dune",
		Short: "dune相关",
		Long:  `获取交易详情`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return getTx(cmd.Context())
		},
	}
	//hash string // 交易hash
)

func init() {

	// Here you will define your flags and configuration settings.

	//txCmd.Flags().StringVarP(&hash, "hash", "h", "", "标签: `m` 助记词")
}

func getDune(ctx context.Context) error {
	// Define manually
	//env = &config.Env{
	//	APIKey: "Your_API_Key",
	//	// you can define a different domain to connect to, for example for a mocked API
	//	Host: "https://api.example.com",
	//}

	// Next, instantiate and use a Dune client object
	//client := dune.NewDuneClient(env)
	//queryID := 1234
	//queryParameters := map[string]any{
	//	"paramKey": "paramValue",
	//}
	//rows, err := client.RunQueryGetRows(queryID, queryParameters)
	//if err != nil {
	//	// handle error
	//}
	//
	//for row := range rows {
	//	// ...
	//}
	return nil
}
