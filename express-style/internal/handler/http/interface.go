package http

type Server interface {
	Run() error
	Shutdown() error
	SetupRoutes()
}
