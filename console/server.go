package console

import (
	"sim-u/config"
	"sim-u/database"
	"sim-u/middleware"
	"sim-u/router"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  "Start running the server",
	Run:   server,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func server(cmd *cobra.Command, args []string) {
	database.InitDB()
	sqlDB, err := database.PostgresDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	e := echo.New()
	e.Use(middleware.LogInfo)
	router.RouterInit(e)
	e.Logger.Fatal(e.Start(":" + config.Port()))
}
