package routes

import (
	"../shared"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"io/ioutil"
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

// TODO Redis
func Tests (c *gin.Context) {
	env := c.MustGet("env").(shared.Env)
	rdb := c.MustGet("rdb").(*redis.ClusterClient)
	var ctx = context.Background()
	tests, _ := rdb.Get(ctx, "tests").Bytes()

	var data pings

	if tests != nil {
		json.Unmarshal(tests, &data)
	} else {
		// TODO helper & FOR TEST Write Happy Case
		url := fmt.Sprintf("%s/tests", env.ApiUrl)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Content-Type", "application/json")

		client := new(http.Client)
		res, _ := client.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &data)
	}

	c.JSON(http.StatusOK, data)
}