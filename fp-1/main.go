package main

import (
	"encoding/json"
	"log"
	"net/http"
	"project-1/database"
	_ "project-1/docs"

	"project-1/controllers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Todo Application
// @version 1.0
// @description This todo application API
// @termsOfService http://swagger.io/terms
// @contact.name API Support
// @contact.email testing@gmail.com
// @license.name Apace 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	database.StartDB()

	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorld).Methods("GET")
	r.HandleFunc("/todos", controllers.GetAllTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", controllers.GetTodoDetail).Methods("GET")
	r.HandleFunc("/todo/{id}", controllers.EditTodo).Methods("PUT")
	r.HandleFunc("/todo/{id}", controllers.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/create-todo", controllers.CreateTodo).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(notfoundHandler)

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Println("Server start at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("page not found"))
}

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode("Hello World")

}
