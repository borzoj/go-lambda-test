package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/borzoj/go-lambda-test/pkg/qna"
	weatherService "github.com/borzoj/go-lambda-test/pkg/weather"
	"github.com/gin-gonic/gin"
)

func foo(c *gin.Context) {
	log.Println("foo")
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "hello lambda API! <-changed"})
}

func ask(c *gin.Context) {
	log.Println("ask")
	var question qna.Question
	err := c.BindJSON(&question)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	response, err := qna.Ask(question)
	c.JSON(http.StatusOK, gin.H{"success": true, "response": response})
}

func weather(c *gin.Context) {
	log.Println("weather")
	city := c.Param("city")
	response, err := weatherService.Get(city)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "weather": response.Weather[0]})
}

func main() {
	addr := ":3000"
	mode := os.Getenv("GIN_MODE")
	g := gin.New()
	g.GET("/foo", foo)
	g.GET("/weather/:city", weather)
	g.POST("/ask", ask)

	if mode == "release" {
		log.Fatal(gateway.ListenAndServe(addr, g))
	} else {
		log.Fatal(http.ListenAndServe(addr, g))
	}
}
