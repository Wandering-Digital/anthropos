package migration

import (
	"log"

	"github.com/Wandering-Digital/anthropos/internal/conn"
	"github.com/Wandering-Digital/anthropos/internal/migration"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the tables in database",
	Long:  `Resets the tables in database`,
	Run:   resetDatabase,
}

func init() {
	RootCmd.AddCommand(resetCmd)
}

func resetDatabase(cmd *cobra.Command, args []string) {
	log.Println("Resetting database...")
	db := conn.GetDB()

	err := db.GormDB.Migrator().DropTable(migration.Models...)
	logDBFatal(err)

	err = db.GormDB.AutoMigrate(migration.Models...)
	logDBFatal(err)

	log.Println("Database Resettled successfully!")
}
