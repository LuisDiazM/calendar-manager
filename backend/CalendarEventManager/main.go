package main

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd"
)

func main() {
	app := cmd.CreateApp()
	app.Start()
}
