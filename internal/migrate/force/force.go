package force

import (
	"errors"
	"github.com/emberfarkas/goctl/internal/migrate/clap"
	"github.com/go-bamboo/pkg/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

var (
	sourceURL   string
	databaseURL string
	version     int
)

var Cmd = &cobra.Command{
	Use:   "force",
	Short: "后退一步数据库",
	Long:  `后退数据库`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cmdParse clap.YamlConfig
		if clap.GenPath != "" {
			configFileParams, err := clap.LoadConfigFile(clap.GenPath)
			if err != nil {
				return err
			}
			cmdParse = *configFileParams
		} else {
			// cmd first
			if sourceURL != "" {
				cmdParse.SourceURL = sourceURL
			}
			if databaseURL != "" {
				cmdParse.DatabaseURL = databaseURL
			}
			if version > 0 {
				cmdParse.Version = version
			}
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
		if err = m.Force(cmdParse.Version); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				return nil
			}
			return err
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&sourceURL, "source", "file://./migrations", "file://./migrations")
	Cmd.Flags().StringVar(&databaseURL, "db", "", "input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]")
	Cmd.Flags().IntVar(&version, "version", -1, "fix and force version")
}
