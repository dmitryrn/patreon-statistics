package internal

import (
	"github.com/pkg/errors"
	controller_patreon_user "patreon-statistics/internal/controller/patreon-user"
	db "patreon-statistics/internal/db"
	http_server "patreon-statistics/internal/http-server"
	repository_patreon_user "patreon-statistics/internal/repository/patreon-user"
	service_patreon_user "patreon-statistics/internal/service/patreon-user"
)

func NewApp() error {
	db, err := db.NewDb()
	if err != nil {
		return errors.Wrap(err, "fail create db")
	}

	patreonUserRepository := repository_patreon_user.NewPatreonUserRepository(db)

	patreonUserService := service_patreon_user.NewPatreonUserService(patreonUserRepository)

	patreonUserController := controller_patreon_user.NewPatreonUserController(patreonUserService)

	httpServer := http_server.NewHttpServer(patreonUserController)

	err = httpServer.Listen()
	if err != nil {
		return errors.Wrap(err, "server listen fail")
	}

	return nil
}
