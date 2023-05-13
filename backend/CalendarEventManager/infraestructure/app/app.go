package app

import (
	"fmt"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
)

type Application struct {
	Env *config.Env
}

func NewApplication(configVars *config.Env) *Application {
	return &Application{
		Env: configVars,
	}
}

func (app *Application) Start() {
	target := fmt.Sprintf(`hi everyone at port %d`, app.Env.PORT)
	fmt.Println(target)
}
