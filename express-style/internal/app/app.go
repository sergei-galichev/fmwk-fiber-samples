package app

import (
	"context"
	"express-style/config"
	"express-style/internal/handler/http"
	"express-style/internal/handler/http/server"
	"github.com/gofiber/fiber/v2/log"
	"github.com/pkg/errors"
)

type App struct {
	provider *ServiceProvider
	server   http.Server
}

func NewApp(ctx context.Context) *App {
	app := &App{}
	err := app.initDeps(ctx)
	if err != nil {
		log.Fatal("App: init dependencies error")
	}

	return app
}

func (a *App) RunApp() {
	err := a.server.Run()
	if err != nil {
		log.Fatal(err)
	}

	a.StopApp()
}

func (a *App) StopApp() {
	log.Info("App: Stopping...")
	err := a.server.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Server: Stopped")

	err = a.provider.Storage().Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Storage: All connections closed")
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initServer,
	}

	for _, fn := range inits {
		if err := fn(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	config.LoadConfig()
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.provider = NewServiceProvider()
	return nil
}

func (a *App) initServer(_ context.Context) error {
	a.server = server.NewServer(a.provider.ProductService())
	a.server.SetupRoutes()
	return nil
}
