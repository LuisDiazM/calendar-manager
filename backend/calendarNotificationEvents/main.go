package main

import (
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/cmd"
)

func main() {
	application := cmd.CreateApp()
	application.Start()
}
