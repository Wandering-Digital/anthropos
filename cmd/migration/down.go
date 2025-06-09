package migration

import (
	"log"

	"github.com/Wandering-Digital/anthropos/internal/conn"
	"github.com/Wandering-Digital/anthropos/internal/migration"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Drop tables from database",
	Long:  `Drop tables from database`,
	Run:   downDatabase,
}

func init() {
	RootCmd.AddCommand(downCmd)
}

func downDatabase(cmd *cobra.Command, args []string) {
	log.Println("Dropping database table...")
	db := conn.GetDB()

	err := db.GormDB.Migrator().DropTable(migration.Models...)
	logDBFatal(err)

	log.Println("Database dopped successfully!")
}
