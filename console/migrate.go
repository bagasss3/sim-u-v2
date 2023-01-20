package console

import (
	"myapp/config"
	"myapp/database"

	"github.com/pressly/goose/v3"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run migrate database",
	Long:  "Start migrate database",
	Run:   migration,
}

func init() {
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction up/down")
	RootCmd.AddCommand(migrateCmd)
}

func migration(cmd *cobra.Command, args []string) {
	direction := cmd.Flag("direction").Value.String()

	err := goose.SetDialect("postgres")
	if err != nil {
		log.Error(err)
	}
	goose.SetTableName("schema_migrations")
	database.InitDB()
	sqlDB, err := database.PostgresDB.DB()
	if err != nil {
		log.WithField("DatabaseDSN", config.DBDSN()).Fatal("Failed to connect database: ", err)
	}
	defer sqlDB.Close()

	var dir string = "./database/migrations"
	if direction == "up" {
		err = goose.Up(sqlDB, dir)
	} else {
		err = goose.Down(sqlDB, dir)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"direction": direction}).
			Fatal("Failed to migrate database: ", err)
	}

	log.WithFields(log.Fields{
		"direction": direction,
	}).Info("Success applied migrations!\n")

}
