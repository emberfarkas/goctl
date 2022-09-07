package benchmark

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/emberfarkas/pkg/client/eth"
	"github.com/ethereum/go-ethereum/ethclient"
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
			return eth.New(url)
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

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
		defer func() {
			wg.Done()
			log.Printf("run done")
			if err := recover(); err != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				pl := fmt.Sprintf("run call panic: %v\n%s\n", err, buf)
				log.Printf("%s", pl)
			}
		}()
		if err := uc.run(ctx); err != nil {
			log.Printf("err: %v", err)
		}
	}(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
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
