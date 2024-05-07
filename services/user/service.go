package user

import (
	"context"

	"encore.app/business/user"
	"encore.app/business/user/db"
	"encore.dev/rlog"
	"encore.dev/storage/sqldb"
)

var appDB = sqldb.Driver(sqldb.Named("app"))

// Service represents the encore service application for users.
//
//encore:service
type Service struct {
	log  rlog.Ctx
	user *user.Business
}

// initService is called by Encore to initialize the service.
//
//lint:ignore U1000 "called by encore"
func initService() (*Service, error) {
	log := rlog.With("service", "user")
	database := db.NewRepository(appDB)

	return &Service{
		log:  log,
		user: user.NewBusiness(log, database),
	}, nil
}

// Shutdown is called by Encore to signal the service that it will be shutdown.
func (s *Service) Shutdown(force context.Context) {
	defer s.log.Info("shutdown", "status", "shutdown complete")
}
