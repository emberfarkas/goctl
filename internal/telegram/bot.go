package telegram

import (
	"context"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "telegram",
	Short: "telegram相关辅助工具",
	Long:  `一些批处理bot的工具`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run1(cmd.Context())
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

func run1(ctx context.Context) error {
	pref := tele.Settings{
		Token:  "5501857852:AAEgPUixc7SD9mVGQSCwMp3_7x_KXqF4nFg",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	var chatId tele.Recipient

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		// b.Send())
		for i := 0; i < 1000; i++ {
			if chatId != nil {
				b.Send(chatId, "小金你好")
			} else {
				time.Sleep(1 * time.Minute)
			}
		}
	}()

	b.Handle("/hello", func(c tele.Context) error {
		chatId = c.Recipient()
		// fmt.Print(chatId)
		return c.Send("Hello!")
	})

	b.Start()
	return nil
}
