package controllers

import "C"
import (
	"encoding/json"
	"github.com/gorilla/mux"
	"http-server/database"
	"http-server/entity"
	"io/ioutil"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// GetAllPicture get all picture data
func GetAllPicture(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var pictures []entity.Picture
	pageSize := 9
	page := r.URL.Query().Get("page")
	pictype := r.URL.Query().Get("type")
	//search := r.URL.Query().Get("search")
	haspage := r.Form["page"]
	haspictype := r.Form["type"]

	pageInt, err := strconv.Atoi(page)
	if err != nil && len(haspage) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid Page Number", http.StatusBadRequest)
	}
	pageInt = (pageInt - 1) * pageSize

	db := database.Connector.Offset(pageInt).Limit(9).Preload("Tags").First(&pictures).Find(&pictures, entity.Picture{Type: pictype})
	errors := db.GetErrors()
	if len(errors) > 0 {
		for i := 0; i < len(errors); i++ {
			s := errors[i].Error()
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, s, http.StatusBadRequest)
		}
	} else {
		if len(haspictype) > 0 {
			switch pictype {
			case "film", "digital":
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(pictures)
			default:
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, "Invalid Query Parameter", http.StatusBadRequest)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pictures)
	}
}

// GetPictureID returns picture with specific ID
func GetPictureID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	key := vars["id"]

	var picture entity.Picture
	db := database.Connector.First(&picture, key)
	errors := db.GetErrors()
	if len(errors) > 0 {
		for i := 0; i < len(errors); i++ {
			s := errors[i].Error()
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, s, http.StatusBadRequest)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(picture)
	}
}

// CreatePicture creates picture
func CreatePicture(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	requestBody, _ := ioutil.ReadAll(r.Body)
	var picture entity.Picture
	json.Unmarshal(requestBody, &picture)

	db := database.Connector.Create(picture)
	errors := db.GetErrors()
	if len(errors) > 0 {
		for i := 0; i < len(errors); i++ {
			s := errors[i].Error()
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, s, http.StatusBadRequest)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(picture)
	}
}

// UpdatePictureByID updates picture with respective ID
func UpdatePictureByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	requestBody, _ := ioutil.ReadAll(r.Body)
	var picture entity.Picture
	json.Unmarshal(requestBody, &picture)

	db := database.Connector.Save(&picture)
	erros := db.GetErrors()
	if len(erros) > 0 {
		for i := 0; i < len(erros); i++ {
			s := erros[i].Error()
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, s, http.StatusBadRequest)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(picture)
	}
}

// DeletePictureByID deletes picture with specific ID
func DeletePictureByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	key := vars["id"]
	var picture entity.Picture
	id, _ := strconv.ParseInt(key, 10, 64)

	db := database.Connector.Where("id = ?", id).Delete(&picture)
	errors := db.GetErrors()
	if len(errors) > 0 {
		for i := 0; i < len(errors); i++ {
			s := errors[i].Error()
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, s, http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
