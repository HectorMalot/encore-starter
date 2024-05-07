package auth

import (
	"context"

	"encore.app/business/auth"
	adb "encore.app/business/auth/db"
	"encore.app/business/user"
	udb "encore.app/business/user/db"
	"encore.dev/rlog"
	"encore.dev/storage/sqldb"
)

var appDB = sqldb.Driver(sqldb.Named("app"))

// Service represents the encore service application for authentication.
//
//encore:service
type Service struct {
	log  rlog.Ctx
	auth *auth.Business
	user *user.Business
}

// initService is called by Encore to initialize the service.
//
//lint:ignore U1000 "called by encore"
func initService() (*Service, error) {
	log := rlog.With("service", "auth")

	return &Service{
		log:  log,
		auth: auth.NewBusiness(log, adb.NewRepository(appDB)),
		user: user.NewBusiness(log, udb.NewRepository(appDB)),
	}, nil
}

// Shutdown is called by Encore to signal the service that it will be shutdown.
func (s *Service) Shutdown(force context.Context) {
	defer s.log.Info("shutdown", "status", "shutdown complete")
}
