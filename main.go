package main

import (
	"net/http"

	"github.com/fercho920/ecommerce-go/config"
	"github.com/fercho920/ecommerce-go/db"
	"github.com/gorilla/mux"
)

func main() {

	db.InitDB()
	if err := db.Migrate(); err != nil {
        panic(err) // Si hay un error en las migraciones, terminar la ejecución del programa
    }



	r := mux.NewRouter()
    config.SetupRoutes(r) 
	http.ListenAndServe(":8080", r)
}
