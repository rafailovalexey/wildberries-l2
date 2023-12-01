package application

import (
	"context"
	"github.com/emptyhopes/wildberries-l2-dev11/cmd/server"
	"github.com/emptyhopes/wildberries-l2-dev11/internal/provider"
	providerEvents "github.com/emptyhopes/wildberries-l2-dev11/internal/provider/events"
	"github.com/joho/godotenv"
)

type InterfaceApplication interface {
	InitializeDependency(ctx context.Context) error
	InitializeEnvironment(_ context.Context) error
	InitializeProvider(_ context.Context) error
	Run()
}

type Application struct {
	providerEvents provider.ProviderEventsInterface
}

var _ InterfaceApplication = (*Application)(nil)

func NewApplication(ctx context.Context) (*Application, error) {
	application := &Application{}

	err := application.InitializeDependency(ctx)

	if err != nil {
		return nil, err
	}

	return application, nil
}

func (a *Application) InitializeDependency(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.InitializeEnvironment,
		a.InitializeProvider,
	}

	for _, function := range inits {
		err := function(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Application) InitializeEnvironment(_ context.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	return nil
}

func (a *Application) InitializeProvider(_ context.Context) error {
	a.providerEvents = providerEvents.NewProviderEvents()

	return nil
}

func (a *Application) Run() {
	api := a.providerEvents.GetControllerEvents()

	server.Run(api)
}
