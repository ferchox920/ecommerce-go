package config

import (
	"github.com/fercho920/ecommerce-go/middleware"
	"github.com/fercho920/ecommerce-go/routes"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", routes.HomeHandler)



	public := r.NewRoute().Subrouter()
	protected:= r.NewRoute().Subrouter()

	protected.Use(middleware.AuthMiddleware)

	public.HandleFunc("/auth/login", routes.LoginHandler).Methods("POST")

	public.HandleFunc("/", routes.HomeHandler)
	public.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	protected.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	protected.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	protected.HandleFunc("/users/update", routes.UpdateUserHandler).Methods("PUT")
	protected.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")


}