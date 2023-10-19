package main

import (
	"http-server/controller"
	"http-server/database"
	"http-server/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/pictures/create", controllers.CreatePicture).Methods("POST")
	router.HandleFunc("/pictures/get", controllers.GetAllPicture).Methods("GET")
	router.HandleFunc("/pictures/get/{id}", controllers.GetPictureID).Methods("GET")
	router.HandleFunc("/pictures/update/{id}", controllers.UpdatePictureByID).Methods("PUT")
	router.HandleFunc("/pictures/delete/{id}", controllers.DeletePictureByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "kotiger",
			Password:   "S@hil007",
			DB:         "pictures",
		}

	connectionString := database.GetConnection(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Picture{}, &entity.Tags{})

}
