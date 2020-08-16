package routes

import (
	"../shared"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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

	// FOR TEST Write Happy Case
	url := fmt.Sprintf("%s/tests", env.ApiUrl)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var data pings
	json.Unmarshal(body, &data)

	c.JSON(http.StatusOK, data)
}