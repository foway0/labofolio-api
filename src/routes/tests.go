package routes

import (
	"../helper"
	"../shared"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

type ping struct {
	Id        int
	Subject   string
	ContentMd string
}

type pings struct {
	Count int
	Rows []ping
}

func Tests (c *gin.Context) {
	env := c.MustGet("env").(shared.Env)
	rdb := c.MustGet("rdb").(*redis.ClusterClient)
	var ctx = context.Background()
	tests, _ := rdb.Get(ctx, "tests").Bytes()

	var data pings

	if tests != nil {
		json.Unmarshal(tests, &data)
	} else {
		url := fmt.Sprintf("%s/tests", env.ApiUrl)
		body := helper.Request("GET", url)
		json.Unmarshal(body, &data)
	}

	c.JSON(http.StatusOK, data)
}