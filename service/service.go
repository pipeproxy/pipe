package service

type Service interface {
	Run() error
	Close() error
}
