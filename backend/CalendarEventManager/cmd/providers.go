package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	meetings "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/usecases"
	users "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/usecases"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
	meetingsRepo "github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database/meetings"
	usersRepo "github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database/users"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server"
	"github.com/google/wire"
)

var EnvironmentVariablesProvider = wire.NewSet(config.NewEnvironmentsSpecification)
var HTTPServerProvider = wire.NewSet(server.NewHTTPServer)
var DatabaseProvider = wire.NewSet(database.NewDatabaseImp)
var AppProvider = wire.NewSet(app.NewApplication)

// usecases
var UsersUsecasesProvider = wire.NewSet(users.NewUsersUsecase)
var MeetingUsecaseProvider = wire.NewSet(meetings.NewMeetingUsecases)

// repositories
var UsersRepositoryProvider = wire.NewSet(usersRepo.NewUsersRepository)
var UsersRepositoryReaderProvider = wire.NewSet(usersRepo.NewUsersRepositoryReader)
var UsersRepositoryUpdateProvider = wire.NewSet(usersRepo.NewUsersRepositoryUpdate)
var MeetingRepositoryCreateProvider = wire.NewSet(meetingsRepo.NewWriteMeetingRepository)
var MeetRepositoryReadProvider = wire.NewSet(meetingsRepo.NewReadMeetingRepo)
