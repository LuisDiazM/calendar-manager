package app

import (
	"context"
	"fmt"
)

func (app *Application) CreateMeetingHandler(ctx context.Context, request map[string]interface{}) error {
	fmt.Println(request)
	return nil
}
