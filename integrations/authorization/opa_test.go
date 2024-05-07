package authorization

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthorization(t *testing.T) {

	policy := `
package encore.rego

import rego.v1

default allow := false

# allow admins to do anything
allow if {
	"admin" in input.user.Roles 
}

# allow users to read their own data
allow if {
	input.user.id == input.resource.user_id
	input.request.Endpoint in {"GetUser", "UpdateUser", "DeleteUser"}
}`

	ctx := context.Background()
	err := evalOpaPolicy(ctx, policy, "allow", map[string]any{
		"user": map[string]any{
			"Roles": []string{"admin", "manager"},
		},
	})
	require.NoError(t, err)
}
