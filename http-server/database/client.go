package database

import (
	"http-server/entity"
	"log"

	"github.com/jinzhu/gorm"
)

// Connector variable used for CRUD operation's
var Connector *gorm.DB

// Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

// Migrate create/updates database table
func Migrate(Picture *entity.Picture, Tags *entity.Tags) {
	/*(var error1 error
	tags := []*entity.Tags{
		{uuid.Must(uuid.NewV4(), error1), "mountain"},   // 1
		{uuid.Must(uuid.NewV4(), error1), "leaves"},     // 2
		{uuid.Must(uuid.NewV4(), error1), "tree"},       // 3
		{uuid.Must(uuid.NewV4(), error1), "oregon"},     // 4
		{uuid.Must(uuid.NewV4(), error1), "2021"},       // 5
		{uuid.Must(uuid.NewV4(), error1), "nature"},     // 6
		{uuid.Must(uuid.NewV4(), error1), "washington"}, // 7
		{uuid.Must(uuid.NewV4(), error1), "river"},      // 8
		{uuid.Must(uuid.NewV4(), error1), "water"},      // 9
		{uuid.Must(uuid.NewV4(), error1), "landscape"},  // 10
	}
	picture := entity.Picture{ID: uuid.Must(uuid.NewV4(), error1), Type: "Digital", Album: "Washington/Oregon 2021", Filename: "S2017763.jpg", Tags: tags}
	Connector.Create(&picture)
	Connector.Save(&picture)*/
	Connector.CreateTable(&Picture)
	Connector.CreateTable(&Tags)
	log.Println("Table migrated")
	/*if error1 != nil {
		fmt.Printf("Something went wrong: %s", error1)
		return
	}*/
}
