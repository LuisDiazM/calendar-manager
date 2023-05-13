package app

import (
	"context"
	"fmt"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	Env       *config.Env
	WebServer *gin.Engine
}

func NewApplication(configVars *config.Env, webServer *gin.Engine) *Application {
	return &Application{
		Env:       configVars,
		WebServer: webServer,
	}
}

func (app *Application) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := app.WebServer.Run(fmt.Sprintf(`:%d`, app.Env.PORT)); err != nil {
			return err
		}
		return nil
	})
	return g.Wait()
}
