package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	users "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/usecases"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
	usersRepo "github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database/repositories/users"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server"
	"github.com/google/wire"
)

var EnvironmentVariablesProvider = wire.NewSet(config.NewEnvironmentsSpecification)
var HTTPServerProvider = wire.NewSet(server.NewHTTPServer)
var DatabaseProvider = wire.NewSet(database.NewDatabaseImp)
var AppProvider = wire.NewSet(app.NewApplication)

// usecases
var UsersUsecasesProvider = wire.NewSet(users.NewUsersUsecase)

// repositories
var UsersRepositoryProvider = wire.NewSet(usersRepo.NewUsersRepository)
var UsersRepositoryReaderProvider = wire.NewSet(usersRepo.NewUsersRepositoryReader)
var UsersRepositoryUpdateProvider = wire.NewSet(usersRepo.NewUsersRepositoryUpdate)
