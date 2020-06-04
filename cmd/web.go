package cmd

import (
	"github.com/Jeiwan/goblogs/db"
	"github.com/Jeiwan/goblogs/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newWebCmd() *cobra.Command {
	var address string
	var dbHost string
	var dbPort int
	var dbName string
	var dbUser string
	var dbPassword string

	cmd := &cobra.Command{
		Use: "web",
		RunE: func(cmd *cobra.Command, args []string) error {
			mainDS := &db.Datasource{
				Host:     dbHost,
				Port:     dbPort,
				DBName:   dbName,
				User:     dbUser,
				Password: dbPassword,
				SSLMode:  false,
			}

			mainDB, err := db.New(mainDS)
			if err != nil {
				return err
			}

			tenantDS := &db.Datasource{
				Host:     dbHost,
				Port:     dbPort,
				DBName:   "tenant",
				User:     dbUser,
				Password: dbPassword,
				SSLMode:  false,
			}

			tenantDB, err := db.New(tenantDS)
			if err != nil {
				return err
			}

			web := web.New(mainDS, mainDB, tenantDB)

			return web.Run(address)
		},
	}

	cmd.Flags().StringVar(&address, "address", viper.GetString("WEB_ADDRESS"), "host:port to listen on")
	cmd.Flags().StringVar(&dbHost, "db-host", viper.GetString("DB_HOST"), "Database host")
	cmd.Flags().IntVar(&dbPort, "db-port", viper.GetInt("DB_PORT"), "Database port")
	cmd.Flags().StringVar(&dbName, "db-name", viper.GetString("DB_NAME"), "Database name")
	cmd.Flags().StringVar(&dbUser, "db-user", viper.GetString("DB_USER"), "Database user")
	cmd.Flags().StringVar(&dbPassword, "db-password", viper.GetString("DB_PASSWORD"), "Database password")

	return cmd
}
