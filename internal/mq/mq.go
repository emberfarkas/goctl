package mq

import (
	"fmt"
	"github.com/emberfarkas/goctl/internal/mq/clap"
	"github.com/go-bamboo/pkg/client/rabbitmq"
	"github.com/go-bamboo/pkg/log"
	"github.com/spf13/cobra"
)

var cfgFile string

var Cmd = &cobra.Command{
	Use:   "mq",
	Short: "mq admin", // 简介
	RunE: func(cmd *cobra.Command, args []string) error {
		var bc *clap.YamlConfig
		bc, err := clap.LoadConfigFile(cfgFile)
		if err != nil {
			return err
		}
		admin := rabbitmq.MustNewAdmin(bc.Conn)
		for _, name := range bc.Names {
			if err := createQueue(bc, admin, name); err != nil {
				log.Error(err)
			}
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVarP(&cfgFile, "conf", "c", "", "configs file")
}

func createQueue(c *clap.YamlConfig, admin *rabbitmq.Admin, name string) error {
	//if err := admin.DeclareQueue(&rabbitmq.AdminQueueConf{
	//	Name:       fmt.Sprintf("queue.%v", name),
	//	Durable:    true,
	//	AutoDelete: false,
	//	Exclusive:  false,
	//	NoWait:     false,
	//}, amqp.Table{
	//	"x-dead-letter-exchange":    fmt.Sprintf("ex.%v", name),
	//	"x-dead-letter-routing-key": "dead",
	//	"x-message-ttl":             c.Ttl.AsDuration().Milliseconds(),
	//}); err != nil {
	//	log.Error(err)
	//}

	//if err := admin.DeclareQueue(&rabbitmq.AdminQueueConf{
	//	Name:       fmt.Sprintf("queue.%v", name),
	//	Durable:    true,
	//	AutoDelete: false,
	//	Exclusive:  false,
	//	NoWait:     false,
	//}, nil); err != nil {
	//	log.Error(err)
	//} else {
	//	log.Infof("declare queue: queue.%v", name)
	//}

	//if err := admin.DeclareQueue(&rabbitmq.AdminQueueConf{
	//	Name:       fmt.Sprintf("queue.%v.newMail", name),
	//	Durable:    true,
	//	AutoDelete: false,
	//	Exclusive:  false,
	//	NoWait:     false,
	//}, nil); err != nil {
	//	log.Error(err)
	//}
	//
	//if err := admin.DeclareQueue(&rabbitmq.AdminQueueConf{
	//	Name:       fmt.Sprintf("queue.%v.recvMail", name),
	//	Durable:    true,
	//	AutoDelete: false,
	//	Exclusive:  false,
	//	NoWait:     false,
	//}, nil); err != nil {
	//	log.Error(err)
	//}

	//if err := admin.DeclareExchange(&rabbitmq.AdminExchangeConf{
	//	Name:       fmt.Sprintf("ex.%v", name),
	//	Kind:       "direct",
	//	Durable:    true, // 持久保存，在rabbitmq宕机或者重启后交换机会不存在。Transient关机就没有了
	//	AutoDelete: false,
	//	Internal:   false,
	//	NoWait:     false,
	//}, nil); err != nil {
	//	log.Error(err)
	//} else {
	//	log.Infof("declare exchange: ex.%v", name)
	//}

	if err := admin.Bind(fmt.Sprintf("queue.%v", name), "normal", fmt.Sprintf("ex.%v", name), rabbitmq.ExchangeDirect, nil); err != nil {
		log.Error(err)
	}

	//if err := admin.Bind(fmt.Sprintf("queue.%v.dead", name), "dead", fmt.Sprintf("ex.%v", name), false, nil); err != nil {
	//	log.Error(err)
	//}

	//if err := admin.Bind(fmt.Sprintf("queue.%v.newMail", name), "newMail", fmt.Sprintf("ex.%v", name), false, nil); err != nil {
	//	log.Error(err)
	//}
	//
	//if err := admin.Bind(fmt.Sprintf("queue.%v.recvMail", name), "recvMail", fmt.Sprintf("ex.%v", name), false, nil); err != nil {
	//	log.Error(err)
	//}
	return nil
}
