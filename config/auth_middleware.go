package config

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"

	"github.com/WilliamsStudentsOnline/wso-backend/models"
)

var identityKey = "id"

type Login struct {
	UnixID string `form:"unix_id" json:"unix_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoadAuthMiddleware(db *gorm.DB) (authMiddleware *jwt.GinJWTMiddleware, err error) {
	// the jwt middleware
	authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			user := new(models.User)
			user.ID = uint(claims["id"].(float64))
			return user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			unixID := loginVals.UnixID
			password := loginVals.Password

			// Do LDAP Authentication HERE
			_ = password

			var user models.User
			// In real version, do FirstOrCreate
			err := db.Where(&models.User{UnixID: unixID}).First(&user).Error
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &user, nil
		},
		Unauthorized: func(c *gin.Context, statusCode int, errorMsg string) {
			c.JSON(statusCode, gin.H{
				"status":    statusCode,
				"error": errorMsg,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	return
}
