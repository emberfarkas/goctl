package benchhalo

import (
	"context"
	"github.com/go-bamboo/pkg/log"
	_ "github.com/go-bamboo/pkg/log/std"
	"github.com/go-bamboo/pkg/threading"
	"github.com/spf13/cobra"
	"sync"
	"time"
)

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:   "benchhalo",
		Short: "压测账号",
		Long:  `压测账号相关的辅助工具`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context())
		},
	}
	flag   string // 标签，判断功能
	priv   string // 私钥
	offset int    // 起始值
	limit  int    // 数量
)

func run(ctx context.Context) error {
	pool, err := threading.NewPool(100)
	if err != nil {
		return err
	}
	defer pool.Close()
	//ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	//n := int32(0)
	////accs := ks.Accounts()
	//w := threading.NewWatch(func(ctx context.Context, account accounts.Account) error {
	//	r := atomic.AddInt32(&n, 1)
	//	log.Infof("-----------------%v", r)
	//	//walletLogin("https://test-api.lifeform.cc/member/pandora/create_user", ks, account)
	//	walletLogin("https://api-v2.lifeform.cc/api/v1/member/wallet_login", ks, account)
	//	return nil
	//}, 10)
	//w.Start()
	//for i := 0; i < 1; i++ {
	//	w.Send(ks.Accounts()[i])
	//}
	//w.Stop()

	//walletLogin("https://api-v2.lifeform.cc/api/v1/member/wallet_login", ks, ks.Accounts()[0])

	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 150; i++ {
		wg.Add(1)
		go walletGen(&wg)
	}
	wg.Wait()
	delta := time.Now().Sub(start)
	log.Infof("%v", delta)
	return nil
}
