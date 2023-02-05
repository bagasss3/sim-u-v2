package console

import (
	"errors"
	"os"
	"os/signal"
	"sim-u/config"
	"sim-u/database"
	"sim-u/worker"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "run worker",
	Long:  "Start running worker",
	Run:   runWorker,
}

func init() {
	RootCmd.AddCommand(workerCmd)
}

func runWorker(_ *cobra.Command, _ []string) {
	// Initiate Connection
	redisConn := database.InitRedis()
	defer redisConn.Close()
	PostgresDB := database.InitDB()
	sqlDB, err := PostgresDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	taskHandler := worker.NewTaskHandler()
	wrk, err := worker.NewWorker(config.RedisWorkerHost(), config.WorkerName(), taskHandler)
	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigChan)
	defer close(sigChan)
	defer close(errChan)

	go func() {
		errChan <- wrk.Run()
	}()
	defer wrk.Stop()

	go func() {
		<-sigChan
		log.Info("Received termination signal, initiating graceful shutdown for workers...")
		errChan <- errors.New("received termination signal")
	}()

	log.Error(<-errChan)
	log.Info("Shutting Down Worker")
}
