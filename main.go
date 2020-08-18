package main

import (
	"./src/middleware"
	"./src/routes"
	"./src/shared"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Init(env shared.Env, config shared.Config, rdb *redis.ClusterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("env", env)
		c.Set("config", config)
		c.Set("rdb", rdb)
	}
}

func main() {
	// init env & config
	env := shared.GetEnv()
	config := shared.GetConfig(env)

	router := gin.Default()

	// TODO config && helper
	var ctx = context.Background()
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			fmt.Sprintf("%s:7000", env.RedisHost),
			fmt.Sprintf("%s:7001", env.RedisHost),
			fmt.Sprintf("%s:7002", env.RedisHost),
			fmt.Sprintf("%s:7003", env.RedisHost),
			fmt.Sprintf("%s:7004", env.RedisHost),
			fmt.Sprintf("%s:7005", env.RedisHost),
		},
	})
	rdb.Ping(ctx)

	// add middleware
	router.Use(Init(env, config, rdb))
	router.Use(middleware.Cors(config.Web.Cors))

	// add routes
	router.GET("/ping", routes.Ping)
	router.GET("/tests", routes.Tests)

	// server start
	router.Run(fmt.Sprintf(":%d", env.ServicePort))
}