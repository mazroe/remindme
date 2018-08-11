package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mazroe/remindme/db"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	db.Save("the world is my country", "MazroE")
	db.Save("Humanity is my race", "MazroE")
	db.Save("Human history is my shame and pride", "MazroE")
	qs, _ := db.All()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"qoutes": qs})
	})

	router.Run(":" + port)
}
