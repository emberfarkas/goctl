package benchmark

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:   "benchmark",
		Short: "压测",
		Long:  `默认1000个账号压测`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context())
		},
	}
	url             string
	mint            string
	contractAddress string
	from            string
	n               int
	chainID         int
	retry           int
	show            int
	pool            = sync.Pool{
		New: func() interface{} {
			c, err := ethclient.Dial(url)
			if err != nil {
				return nil
			}
			return c
		},
	}
	rpcpool = sync.Pool{
		New: func() interface{} {
			c, err := ethclient.Dial(url)
			if err != nil {
				return nil
			}
			return c
		},
	}
)

var (
	errInvalidContractAddress = errors.New("invalid contract address")
	errInvalidChainID         = errors.New("invalid chain ID")
	errNotPubECDSA            = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
)

func init() {

	// Here you will define your flags and configuration settings.
	Cmd.Flags().StringVarP(&url, "url", "u", "", "url")
	Cmd.Flags().StringVarP(&mint, "mint", "m", "mint", "标记")
	Cmd.Flags().StringVar(&from, "from", "", "有钱地址")
	Cmd.Flags().StringVarP(&contractAddress, "contract", "x", "", "合约地址")
	Cmd.Flags().IntVarP(&n, "n", "n", 1, "压测次数")
	Cmd.Flags().IntVar(&chainID, "chainid", 7210, "链ID")
	Cmd.Flags().IntVar(&retry, "retry", 0, "重试")
	Cmd.Flags().IntVar(&show, "show", 0, "显示查询结果")
}

func run(ctx context.Context) error {
	uc, err := newBiz()
	if err != nil {
		return err
	}

	// 开始启动
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		rescue.Recover(func() {
			wg.Done()
			log.Info("run done")
		})
		if err := uc.run(ctx); err != nil {
			log.Errorf("err: %v", err)
		}
	}(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cancel()
			wg.Wait()
			if err = uc.close(); err != nil {
				return err
			}
			time.Sleep(time.Second)
			return nil
		case syscall.SIGHUP:
		default:
			return nil
		}
	}
}
