package worker

import (
	"sim-u/model"
	"time"

	"github.com/hibiken/asynq"
	log "github.com/sirupsen/logrus"
)

var mux = asynq.NewServeMux()

type Worker struct {
	server      *asynq.Server
	client      *asynq.Client
	taskHandler *TaskHandler
	name        string
}

func NewWorker(redis, name string, taskHandler *TaskHandler) (model.WorkerController, error) {
	redisOpt, err := asynq.ParseRedisURI(redis)
	if err != nil {
		return nil, err
	}

	client := asynq.NewClient(redisOpt)

	workerSrv := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency:         10,
		Logger:              log.WithField("worker-name", name),
		HealthCheckFunc:     healthCheck,
		HealthCheckInterval: 5 * time.Minute,
	})

	return &Worker{
		server:      workerSrv,
		client:      client,
		taskHandler: taskHandler,
		name:        name,
	}, nil
}

func (aw *Worker) Run() error {
	if err := aw.server.Run(mux); err != nil {
		log.Fatalf("could not run worker server: %v", err)
		return err
	}
	return nil
}

func (aw *Worker) Stop() {
	defer aw.server.Shutdown()
	if aw.client != nil {
		err := aw.client.Close()
		log.Fatalf("error close client worker: %v", err)
	}
	if aw.server != nil {
		aw.server.Stop()
	}
}

func healthCheck(err error) {
	if err != nil {
		log.Errorf("worker error: %v", err)
	}
}
