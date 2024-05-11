package app

import (
	"context"
	"fmt"
	"log"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	meetingEntities "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	userEntities "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"

	meetings "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/usecases"
	users "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/usecases"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/messaging"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	Env             *config.Env
	WebServer       *gin.Engine
	Database        database.DatabaseGateway
	UsersUsecase    *users.UsersUsecase
	MeetingsUsecase *meetings.MeetingsUsecase
	BrokerProducer  *messaging.RabbitProducer
}

func NewApplication(configVars *config.Env,
	webServer *gin.Engine,
	database database.DatabaseGateway,
	usersUsecase *users.UsersUsecase,
	meetingUsecase *meetings.MeetingsUsecase,
	brokerProducer *messaging.RabbitProducer,
) *Application {
	return &Application{
		Env:             configVars,
		WebServer:       webServer,
		Database:        database,
		UsersUsecase:    usersUsecase,
		MeetingsUsecase: meetingUsecase,
		BrokerProducer:  brokerProducer,
	}
}

func (app *Application) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		app.Database.SetUp(*app.Env)
		//app.Migrations()
		if err := app.WebServer.Run(fmt.Sprintf(`:%d`, app.Env.SERVER_PORT)); err != nil {
			return err
		}
		log.Println("application started!!")
		return nil
	})
	return g.Wait()
}

func (app *Application) Migrations() {
	app.Database.DB().AutoMigrate(&meetingEntities.Meetings{}, &userEntities.Users{})
}
