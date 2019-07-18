package config

import (
	"errors"
	"fmt"
	"github.com/aidanlloydtucker/wso-backend-go-demo/controllers"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ScopeAdminAll      = "admin:all"
	ScopeReadAll       = "read:all"
	ScopeWriteSelf     = "write:self"
	ScopeReadEphcatch  = "read:ephcatch"
	ScopeWriteEphcatch = "write:ephcatch"
	ScopeAdminFactrak  = "admin:factrak"
)

func RequireScopes(scopes ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		fmt.Println(claims)
		jwtScopesIface := (claims["scope"]).([]interface{})
		jwtScopes := make([]string, len(jwtScopesIface))
		for i, v := range jwtScopesIface {
			jwtScopes[i] = v.(string)
		}

		authed := false
		for _, scope := range scopes {
			if authed = containsScope(jwtScopes, scope); authed {
				break
			}
		}

		if !authed {
			controllers.Base.RespondError(
				http.StatusForbidden, errors.New("user does not have scope authorization"), c)
			c.Abort()
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
