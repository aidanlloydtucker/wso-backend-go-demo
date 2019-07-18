package main

import (
	"errors"
	"github.com/aidanlloydtucker/wso-backend-go-demo/config"
	"github.com/aidanlloydtucker/wso-backend-go-demo/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/fvbock/endless"

	"github.com/aidanlloydtucker/wso-backend-go-demo/controllers"
)

func main() {
	/* DATABASE */
	db := config.LoadDatabase()
	defer config.CloseDatabase(db)

	/* SERVER */
	r := gin.New()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Build JWT auth middleware
	authMiddleware, err := config.LoadAuthMiddleware(db)
	if err != nil {
		log.Fatalln("JWT Error: " + err.Error())
	}

	/* ROUTER */

	// Initialize login
	r.POST("/api/v1/auth/login", authMiddleware.LoginHandler)

	// Require authentication for 404s
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		controllers.Base.RespondError(http.StatusNotFound, errors.New("page not found"), c)
	})

	// Wrap everything else in authentication
	router := r.Group("")
	router.Use(authMiddleware.MiddlewareFunc())

	// Middleware to set the user and user id with the context
	router.Use(func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		userID := uint(claims["id"].(float64))
		user := models.NewUserWithID(userID)
		c.Set("user", &user)
		c.Set("user_id", userID)
		c.Next()
	})

	// Actual API routing
	v1 := router.Group("/api/v1")
	{
		v1.GET("/auth/refresh_token", authMiddleware.RefreshHandler)

		userGroup := v1.Group("/user")
		userControl := controllers.NewUserController(db)
		userGroup.GET("/", userControl.FetchAllUsers)
		userGroup.GET("/:user_id", userControl.GetUser)
		userGroup.PUT("/:user_id", userControl.UpdateUser)
	}

	// Would change this to be more production-friendly in real life. I'd use something like endless to keep
	// the server running even when it crashes
	endless.ListenAndServe(":8080", r) // listen and serve on 0.0.0.0:8080
}
