package main

import (
	"net/http"

	"github.com/Psychohistorian11/golang-restAPI-PostgreSQL/db"
	"github.com/Psychohistorian11/golang-restAPI-PostgreSQL/models"
	"github.com/Psychohistorian11/golang-restAPI-PostgreSQL/routes"
	"github.com/gorilla/mux"
)

func main() {
	//Incialización base de datos
	db.DBConnection()

	//Creación de tablas
	db.DB.AutoMigrate(&models.User{}, &models.Task{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//tasks routers
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
