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
	Id        int `json:"id"`
	Subject   string `json:"subject"`
	ContentMd string `json:"content_md"`
}

type pings struct {
	Count int `json:"count"`
	Rows []ping `json:"rows"`
}

func Tests (c *gin.Context) {
	env := c.MustGet("env").(shared.Env)
	rdb := c.MustGet("rdb").(*redis.ClusterClient)
	var ctx = context.Background()
	rdb.Del(ctx, "tests")
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