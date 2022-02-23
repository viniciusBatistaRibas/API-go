package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/viniciusBatistaRibas/go-rest-api.git/controllers"
	"github.com/viniciusBatistaRibas/go-rest-api.git/middleware"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.Use(middleware.Content)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonas).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.Retorna_personalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.Criar).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.Deleta).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.Edita).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
