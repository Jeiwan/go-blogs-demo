package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Run() error {
	viper.AutomaticEnv()

	rootCmd := &cobra.Command{
		Use: "cmd",
	}

	rootCmd.AddCommand(newUICmd())
	rootCmd.AddCommand(newWebCmd())

	return rootCmd.Execute()
}
