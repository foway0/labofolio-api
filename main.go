package main

import (
	"./src/middleware"
	"./src/routes"
	"./src/shared"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init(env shared.Env, config shared.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("env", env)
		c.Set("config", config)
	}
}

func main() {
	// init env & config
	env := shared.GetEnv()
	config := shared.GetConfig(env)

	router := gin.Default()

	// add middleware
	router.Use(Init(env, config))
	router.Use(middleware.Cors(config.Web.Cors))

	// add routes
	router.GET("/ping", routes.Ping)
	router.GET("/tests", routes.Tests)

	// server start
	router.Run(fmt.Sprintf(":%d", env.ServicePort))
}