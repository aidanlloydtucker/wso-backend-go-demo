package main

import (
	"github.com/aidanlloydtucker/wso-backend-go-demo/config"
	"github.com/aidanlloydtucker/wso-backend-go-demo/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/aidanlloydtucker/wso-backend-go-demo/controllers"
)

func main() {
	/* DATABASE */
	db := config.LoadDatabase()
	defer config.CloseDatabase(db)

	/* SERVER */
	r := gin.New()
	r.Use(gin.Logger())
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
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Page not found"})
	})

	// Wrap everything else in authentication
	router := r.Group("")
	router.Use(authMiddleware.MiddlewareFunc())

	// Middleware to set the user and user id with the context
	router.Use(func (c *gin.Context) {
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
	}

	// Would change this to be more production-friendly in real life. I'd use something like endless to keep
	// the server running even when it crashes
	r.Run() // listen and serve on 0.0.0.0:8080
}