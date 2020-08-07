package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// FIXME import
type ping struct {
	Id        int
	Subject   string
	ContentMd string
}
type pings struct {
	Count int
	Rows []ping
}

func main() {
	router := gin.Default()

	// FIXME import
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!!")
	})

	// FIXME import
	router.GET("/ping", func(c *gin.Context) {
		// FOR TEST Write Happy Case
		url := "http://local-api:3000/tests"
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Content-Type", "application/json")

		client := new(http.Client)
		res, _ := client.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		var data pings
		json.Unmarshal(body, &data)

		// FIXME cors middlewares
		c.Header("Access-Control-Allow-Origin", "https://localhost:8080")
		c.Header("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.JSON(http.StatusOK, data)
	})

	// FIXME environment
	router.Run(":3000")
}