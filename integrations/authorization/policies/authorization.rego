package encore.rego

import rego.v1

default allow := false

# allow admins to do anything
allow if {
    "admin" in input.user.Roles 
}

# allow users to read, update and delete their own profile
allow if {
    input.user.id == input.resource.id
    input.request.Service == "user"
    input.request.Endpoint in {"GetUser", "UpdateUser", "DeleteUser"}
}

# See https://www.openpolicyagent.org/docs/latest/policy-language/
# for more information about the Rego language
