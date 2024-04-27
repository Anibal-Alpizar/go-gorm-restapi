package routes

import (
	"encoding/json"
	"net/http"

	"github.com/anibal-alpizar/go-gorm-fastapi/db"
	"github.com/anibal-alpizar/go-gorm-fastapi/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}
	json.NewEncoder(w).Encode(task)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	db.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}
	// db.DB.Delete(&task) // (logical delete)
	db.DB.Unscoped().Delete(&task) // (physical delete)
	w.WriteHeader(http.StatusNoContent) // 204
}
