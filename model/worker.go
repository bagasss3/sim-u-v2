package model

type WorkerController interface {
	Run() error
	Stop()
}
