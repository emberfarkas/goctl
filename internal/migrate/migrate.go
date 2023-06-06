package migrate

import (
	"github.com/emberfarkas/goctl/internal/migrate/clap"
	"github.com/emberfarkas/goctl/internal/migrate/new"
	"github.com/emberfarkas/goctl/internal/migrate/up"

	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source"

	"github.com/spf13/cobra"
)

// Cmd represents the new command
var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate模块",
	Long:  `migrate相关的统计`,
}

func init() {
	Cmd.AddCommand(new.Cmd, up.Cmd)
	Cmd.PersistentFlags().StringVar(&clap.GenPath, "c", "", "is path for gen.yml")
}
