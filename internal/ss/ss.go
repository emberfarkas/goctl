package ss

import (
	"context"
	"github.com/go-bamboo/pkg/log"
	"github.com/imroc/req/v3"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:  "ss",
	Long: `一些批处理proxy的工具`,
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
	//socks5proxy, err := proxy.SOCKS5("tcp", "127.0.0.1:1079", nil, &net.Dialer{
	//	Timeout:   30 * time.Second,
	//	KeepAlive: 30 * time.Second,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//transport := &stdhttp.Transport{
	//	Proxy:               nil,
	//	Dial:                socks5proxy.Dial,
	//	TLSHandshakeTimeout: 10 * time.Second,
	//}
	body, err := req.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%v", body)
	return nil
}
