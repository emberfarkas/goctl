package migrate

import (
	"github.com/emberfarkas/goctl/internal/migrate/clap"
	"github.com/emberfarkas/goctl/internal/migrate/down"
	"github.com/emberfarkas/goctl/internal/migrate/force"
	"github.com/emberfarkas/goctl/internal/migrate/new"
	"github.com/emberfarkas/goctl/internal/migrate/up"
	"github.com/spf13/cobra"

	// clickhouse
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	// mysql
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	// postgres
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// sqlite3
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/source"
)

// Cmd represents the new command
var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate模块",
	Long:  `migrate相关的统计`,
}

func init() {
	Cmd.AddCommand(new.Cmd, up.Cmd, down.Cmd, force.Cmd)
	Cmd.PersistentFlags().StringVarP(&clap.GenPath, "conf", "c", "", "is path for gen.yml")
}
