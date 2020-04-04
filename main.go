package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	// TODO: use env
	// TODO: enable ssl
	db, err = gorm.Open("postgres", "host=localhost port=5454 user=postgres dbname=automobile_api password=pw sslmode=disable")
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
