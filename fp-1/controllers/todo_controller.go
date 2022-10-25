package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project-1/database"
	"project-1/models"
	"time"

	"github.com/gorilla/mux"
)

// GetAllTodos godoc
// @Summary Get all todos
// @Description Get details of all To Do list
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func GetAllTodos(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var todos []models.Todo

	db := database.GetDB()

	db.Find(&todos)
	json.NewEncoder(rw).Encode(todos)
}

// GetTodoDetail godoc
// @Summary Get todo detail
// @Description Get details of todo by id
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "ID of the todo"
// @Success 200 {object} models.Todo
// @Router /todo/{id} [get]
func GetTodoDetail(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var (
		vars = mux.Vars(r)
		id   = vars["id"]
		todo models.Todo
	)

	fmt.Println(id)

	db := database.GetDB()

	err := db.First(&todo, id).Error
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{Message: "Data not found"})
		return
	}

	json.NewEncoder(rw).Encode(todo)
}

// CreateTodo godoc
// @Summary Create todo
// @Description create todo from parameter
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.TodoPayload true "Create todo"
// @Success 200 {object} models.Response
// @Router /create-todo [post]
func CreateTodo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var todo models.Todo
	var payload models.TodoPayload

	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(c, &payload); err != nil {
		panic(err)
	}

	todo.CreatedAt = time.Now()
	todo.Description = payload.Description
	fmt.Println(todo)

	db := database.GetDB()

	err = db.Create(&todo).Error
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(models.Response{Message: "Create failed"})
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(models.Response{Data: todo, Message: "Data created."})
}

// EditTodo godoc
// @Summary Edit todo indentified by the given todo id
// @Description Edit todo using the todo id
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "ID of the todo to be edited"
// @Param todo body models.TodoPayload true "Todo description"
// @Success 200 {object} models.Response
// @Router /todo/{id} [put]
func EditTodo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var (
		vars    = mux.Vars(r)
		id      = vars["id"]
		todo    models.Todo
		payload models.TodoPayload
	)

	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(c, &payload); err != nil {
		panic(err)
	}
	db := database.GetDB()

	err = db.First(&todo, id).Error
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{Message: "Data not found"})
		return
	}

	todo.UpdatedAt = time.Now()
	todo.Description = payload.Description

	err = db.Save(&todo).Error
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(models.Response{Message: "Update failed"})
		return
	}

	json.NewEncoder(rw).Encode(models.Response{Data: todo, Message: "Update Success"})
}

// DeleteTodo godoc
// @Summary Delete todo identified by the given id
// @Description Delete todo corresponding to the input id
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the todo to be deleted"
// @Success 204 {object} models.Response
// @Router /todo/{id} [delete]
func DeleteTodo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var (
		vars = mux.Vars(r)
		id   = vars["id"]
		todo models.Todo
	)

	db := database.GetDB()

	err := db.Delete(&todo, id).Error
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{Message: "Delete fail, error :" + err.Error()})
		return
	}

	json.NewEncoder(rw).Encode(models.Response{Data: todo, Message: "Delete Success"})
}
