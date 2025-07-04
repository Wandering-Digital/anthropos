package migration

import (
	"fmt"
	"log"

	"github.com/Wandering-Digital/anthropos/internal/conn"

	"github.com/spf13/cobra"
)

var (

	// RootCmd represents the base command when called without any subcommands
	RootCmd = &cobra.Command{
		Use:   "migration",
		Short: "Run database migrations",
		Long:  `Migration is a tool to generate and modify databse tables`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := conn.ConnectDB(); err != nil {
				return fmt.Errorf("Can't connect database: %v", err)
			}
			return nil
		},
	}
)

func logDBFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}
