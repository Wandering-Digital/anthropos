package cmd

import (
	"log"

	"github.com/Wandering-Digital/anthropos/cmd/migration"
	"github.com/Wandering-Digital/anthropos/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// cfgFile store the configuration file name
	cfgFile                 string
	verbose, prettyPrintLog bool

	// rootCmd is the root command of backup service
	rootCmd = &cobra.Command{
		Use:   "anthropos",
		Short: "anthropos is a basic golang backend service",
		Long:  `anthropos basic golang service based on clean code architecture`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file")
	rootCmd.PersistentFlags().BoolVarP(&prettyPrintLog, "pretty", "p", false, "pretty print verbose/log")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// set the value to viper config
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// add migration root
	rootCmd.AddCommand(migration.RootCmd)
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func initConfig() {
	log.Println("Loading configurations")
	config.Init()
	log.Println("Configurations loaded successfully!")
}
