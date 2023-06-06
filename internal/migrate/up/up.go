package up

import (
	"github.com/emberfarkas/goctl/internal/migrate/clap"
	"github.com/go-bamboo/pkg/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

var (
	sourceURL   string
	path        string
	databaseURL string
)

var Cmd = &cobra.Command{
	Use:   "up",
	Short: "更新数据库",
	Long:  `更新数据库`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cmdParse clap.YamlConfig
		if clap.GenPath != "" {
			if configFileParams, err := clap.LoadConfigFile(clap.GenPath); err == nil && configFileParams != nil {
				cmdParse = *configFileParams
			}
		}
		// cmd first
		if sourceURL != "" {
			cmdParse.SourceURL = sourceURL
		}
		if sourceURL == "" && path != "" {
			cmdParse.SourceURL = path
		}
		if databaseURL != "" {
			cmdParse.DatabaseURL = databaseURL
		}
		m, err := migrate.New(cmdParse.SourceURL, cmdParse.DatabaseURL)
		if err != nil {
			return err
		}
		defer func() {
			if _, err := m.Close(); err != nil {
				log.Error(err)
			}
		}()
		if err = m.Up(); err != nil {
			if err.Error() == "no change" {
			}
			return err
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&sourceURL, "source", "file://./migrations", "file://./migrations")
	Cmd.Flags().StringVar(&path, "path", "./migrations", "specify a directory for output")
	Cmd.Flags().StringVar(&databaseURL, "db", "", "input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]")
}
