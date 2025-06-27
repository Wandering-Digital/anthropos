package migration

import (
	"log"

	"github.com/Wandering-Digital/anthropos/internal/conn"
	"github.com/Wandering-Digital/anthropos/internal/migration"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Populate tables in database",
	Long:  `Populate tables in database`,
	Run:   upDatabase,
}

func init() {
	RootCmd.AddCommand(upCmd)
}

func upDatabase(cmd *cobra.Command, args []string) {
	log.Println("Populating database...")
	db := conn.GetDB()

	err := db.GormDB.AutoMigrate(migration.Models...)
	logDBFatal(err)

	log.Println("Database populated successfully!")
}
