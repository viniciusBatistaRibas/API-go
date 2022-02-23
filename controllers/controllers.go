package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viniciusBatistaRibas/go-rest-api.git/database"
	"github.com/viniciusBatistaRibas/go-rest-api.git/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")

}

func TodasPersonas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","Aplication/json")
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func Retorna_personalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func Criar(w http.ResponseWriter, r *http.Request) {
	var persona models.Personalidade
	json.NewDecoder(r.Body).Decode(&persona)
	database.DB.Create(&persona)
	json.NewEncoder(w).Encode(persona)

}

func Deleta(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade,id)
	json.NewEncoder(w).Encode(personalidade)
}

func Edita(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
