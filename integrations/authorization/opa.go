package authorization

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"encore.dev/rlog"
	"github.com/open-policy-agent/opa/rego"
)

const opaPackage string = "encore.rego"

//go:embed policies/authorization.rego
var opaAuthorization string

// Authorize attempts to authorize the user with the provided input data.
// If the user is authorized, the function returns nil, otherwise an error is returned.
func Authorize(ctx context.Context, input map[string]any) error {
	return evalOpaPolicy(ctx, opaAuthorization, "allow", input)
}

// evalOpaPolicy evaluates the provided rego policy with the input data
// and returns an error if the policy evaluation fails.
func evalOpaPolicy(ctx context.Context, policy, rule string, input map[string]any) error {
	query, err := rego.New(
		rego.Query(fmt.Sprintf("data.%s.%s", opaPackage, rule)),
		rego.Module("policy.rego", policy),
	).PrepareForEval(ctx)
	if err != nil {
		return err
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))
	rlog.Info(fmt.Sprintf("result: %+v\nerr: %+v\n", results, err))
	if err != nil || !results.Allowed() {
		return errors.New("policy evaluation failed")
	}

	return nil
}
