package leveldb

import (
	"context"

	"github.com/go-bamboo/pkg/log"
	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "leveldb",
	Short: "leveldb相关辅助工具",
	Long:  `一些批处理leveldb的工具`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func run(ctx context.Context) error {
	// 恢复不了metamask数据
	opts := &opt.Options{}
	db, err := leveldb.OpenFile("D:/dev/extensions1/D", opts)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var stats leveldb.DBStats
	if err = db.Stats(&stats); err != nil {
		panic(err)
	}

	log.Info("level size:", stats.LevelSizes, len(stats.LevelSizes))
	log.Info("level tables counts:", stats.LevelTablesCounts, len(stats.LevelDurations))
	return nil
}
