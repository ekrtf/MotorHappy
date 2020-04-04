package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type automobile struct {
	AutomobileID uuid.UUID `json:"automobile_id" gorm:"type:uuid;primary_key;default:uuid_generate_v4();"`
	MakeName     string    `json:"make_name"`
	ModelName    string    `json:"model_name"`
	Nationality  string    `json:"nationality"`
	Version      string    `json:"version"`
	Horsepower   int       `json:"horsepower"`
}

func init() {
	var err error
	// TODO: use env
	// TODO: enable ssl
	db, err = gorm.Open("postgres", "host=localhost port=5454 user=postgres dbname=automobile_api password=pw sslmode=disable")
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&automobile{})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/v1/automobiles", getAutomobiles)

	r.Run(":8080")
}

func getAutomobiles(c *gin.Context) {
	var autos []automobile

	db.Find(&autos)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": autos})
}
