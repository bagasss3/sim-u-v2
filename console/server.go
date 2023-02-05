package console

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sim-u/config"
	"sim-u/controller"
	"sim-u/database"
	"sim-u/middleware"
	"sim-u/repository"
	"sim-u/router"
	"sim-u/service"
	"syscall"
	"time"

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
	// Initiate Connection
	redisConn := database.InitRedis()
	defer redisConn.Close()
	PostgresDB := database.InitDB()
	sqlDB, err := PostgresDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Create Echo instance
	httpServer := echo.New()

	// Initiate Depedency
	studentRepository := repository.NewStudentRepository(PostgresDB)
	studentController := controller.NewStudentController(studentRepository)
	studentService := service.NewStudentService(studentController)
	httpServer.Use(middleware.LogInfo)
	router.RouteService(httpServer.Group(""), studentService)

	// Graceful Shutdown
	// Catch Signal
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigChan)
		defer close(sigChan)

		<-sigChan
		log.Info("Received termination signal, initiating graceful shutdown...")
		cancel()
	}()

	// Start http server
	go func() {
		log.Info("Starting server...")
		if err := httpServer.Start(":" + config.Port()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting the server: %v", err)
		}
	}()

	// Shutting down any connection and server
	<-ctx.Done()
	log.Info("Shutting down server...")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Info("Server gracefully shut down")
}
