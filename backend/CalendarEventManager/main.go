package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server/routes"
)

func main() {
	app := cmd.CreateApp()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	routes.SetUpRoutes(app)
	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
