package config

import (
	"errors"
	"net/http"

	"github.com/aidanlloydtucker/wso-backend-go-demo/controllers"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// The current scopes
const (
	ScopeAdminAll      = "admin:all"
	ScopeReadAll       = "read:all"
	ScopeWriteSelf     = "write:self"
	ScopeReadEphcatch  = "read:ephcatch"
	ScopeWriteEphcatch = "write:ephcatch"
	ScopeAdminFactrak  = "admin:factrak"
)

// Require this endpoint to have a scope; multiple scopes mean an OR. For an AND, call this function multiple times
func RequireScopes(scopes ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)

		// Extract the scope
		jwtScopesIface := (claims["scope"]).([]interface{})
		jwtScopes := make([]string, len(jwtScopesIface))
		for i, v := range jwtScopesIface {
			jwtScopes[i] = v.(string)
		}

		// Check if the scope is valid
		authed := false
		for _, scope := range scopes {
			if authed = containsScope(jwtScopes, scope); authed {
				break
			}
		}

		// If it isn't, abort with error
		if !authed {
			controllers.Base.RespondError(
				http.StatusForbidden, errors.New("user does not have scope authorization"), c)
			return
		}

		c.Next()
	}
}

func containsScope(slice []string, scope string) bool {
	for _, elem := range slice {
		if elem == scope {
			return true
		}
	}
	return false
}
