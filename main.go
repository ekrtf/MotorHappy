package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type make struct {
	MakeID      uint   `json:"make_id" gorm:"primary_key"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
}

type vehiculeType struct {
	VehiculeTypeID uint   `json:"vehicule_type_id" gorm:"primary_key"`
	Name           string `json:"name"`
}

type vehicule struct {
	VehiculeID     uint   `json:"vehicule_id" gorm:"primary_key"`
	VehiculeTypeID uint   `json:"vehicle_type_id"` // TODO: foreign key
	MakeID         uint   `json:"make_id"`         // TODO: foreigh key
	ModelName      string `json:"model_name"`
}

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

	db.AutoMigrate(&make{}, &vehiculeType{}, &vehicule{})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
