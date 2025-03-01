package app

import (
	"context"
	"github.com/auth-service/internal/config"
	"github.com/auth-service/internal/initialize"
	"github.com/auth-service/pkg"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"time"
)

const shutdownTimeout = 15 * time.Second

func Run(config *config.Config) error {
	ctx := context.Background()

	dbConnect, err := pgx.Connect(ctx, config.GetDatabaseConnect())
	if err != nil {
		return errors.Wrap(err, "failed to connect to database")
	}

	repositories := initialize.NewRepositories(dbConnect)
	services := initialize.NewServices(repositories)
	controllers := initialize.NewControllers(services)
	routes := initialize.NewRoutes(controllers)

	server := new(pkg.Server)
	if err := server.Run(config.HTTP_PORT, routes.InitRoutes()); err != nil {
		return errors.Wrap(err, "failed to start server")
	}

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		return errors.Wrap(err, "failed to shutdown server")
	}

	return nil
}
