package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type make struct {
	MakeID      uuid.UUID `json:"make_id" gorm:"type:uuid;primary_key;default:uuid_generate_v4();"`
	Name        string    `json:"name"`
	Nationality string    `json:"nationality"`
}

type vehiculeType struct {
	VehiculeTypeID uuid.UUID `json:"vehicule_type_id" gorm:"type:uuid;primary_key;default:uuid_generate_v4();"`
	Name           string    `json:"name"`
}

type vehicule struct {
	VehiculeID     uuid.UUID    `json:"vehicule_id" gorm:"type:uuid;primary_key;default:uuid_generate_v4();"`
	vehiculeType   vehiculeType `gorm:"foreignkey:VehiculeTypeID"`
	VehiculeTypeID uuid.UUID    `json:"vehicle_type_id"`
	make           uuid.UUID    `gorm:"foreignkey:MakeID"`
	MakeID         uuid.UUID    `json:"make_id"`
	ModelName      string       `json:"model_name"`
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
