package console

import (
	"sim-u/config"
	"sim-u/controller"
	"sim-u/database"
	"sim-u/middleware"
	"sim-u/repository"
	"sim-u/router"
	"sim-u/service"

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
	studentRepository := repository.NewStudentRepository(database.PostgresDB)
	studentController := controller.NewStudentController(studentRepository)
	studentService := service.NewStudentService(studentController)
	e.Use(middleware.LogInfo)
	router.RouteService(e.Group(""), studentService)
	e.Logger.Fatal(e.Start(":" + config.Port()))
}
