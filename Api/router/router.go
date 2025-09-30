package router

import (
	"cenox/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", handlers.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users", handlers.GetUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods(http.MethodDelete)

	return router
}
