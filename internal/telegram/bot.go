package telegram

import (
	"context"

	"github.com/go-bamboo/pkg/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "telegram",
	Short: "telegram相关辅助工具",
	Long:  `一些批处理bot的工具`,
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
	bot, err := tgbotapi.NewBotAPI("5501857852:AAEgPUixc7SD9mVGQSCwMp3_7x_KXqF4nFg")
	if err != nil {
		return err
	}
	for i := 0; ; i++ {
		u := tgbotapi.NewUpdate(i)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message != nil { // If we got a message
				log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
