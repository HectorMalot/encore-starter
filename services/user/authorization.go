package user

import (
	"encore.app/integrations/authorization"
	"encore.app/services/auth"
	"encore.dev/beta/errs"
	"encore.dev/middleware"
	"encore.dev/types/uuid"
)

// Authorize attempts to authorize the user with the provided input data
// This middleware runs for all endpoints with 'tag:authorize' within the user service
//
//lint:ignore U1000 "called by encore"
//encore:middleware target=tag:authorize
func (s *Service) Authorize(req middleware.Request, next middleware.Next) middleware.Response {
	cu, err := auth.CurrentUser()
	if err != nil {
		s.log.Error("Authorize middleware called with invalid current user", "error", err)
		return middleware.Response{Err: &errs.Error{
			Code:    errs.PermissionDenied,
			Message: "Unauthorized",
		}}
	}

	input := map[string]any{
		"user":    cu,
		"request": req.Data(),
	}

	// Extract the ID of the requested user, if it exists
	if len(req.Data().PathParams) == 1 && req.Data().PathParams[0].Name == "id" {
		id, err := uuid.FromString(req.Data().PathParams[0].Value)
		if err != nil {
			s.log.Error("Authorize middleware expected a user ID in the request path but found none", "path_params", req.Data().PathParams, "error", err)
		} else {
			// Add the user ID to the input map
			input["resource"] = map[string]any{"id": id}

			// Note: This is a good place to fetch an entity from the database and add it to the input map
			// For example:
			// product, err := s.db.GetProductByID(id)
			// ...
			// input["resource"] = product
			//
			// It so happens that for the user service, the user_id is already available in the path params
			// and we don't need to fetch the user from the database to deterime ownership
		}
	}

	s.log.Debug("Authorize middleware called", "input", input)

	// Check request agains the authorization policy
	err = authorization.Authorize(req.Context(), input)
	if err != nil {
		return middleware.Response{Err: &errs.Error{
			Code:    errs.PermissionDenied,
			Message: "Unauthorized",
		}}
	}

	return next(req)
}
