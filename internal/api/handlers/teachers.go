package handlers

import (
	"academy-app-system/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

// initialize some dummy data - in memory data
func init() {
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextID++
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Joffery",
		LastName:  "Baratheon",
		Class:     "10B",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Joffery",
		LastName:  "Doe",
		Class:     "11C",
		Subject:   "Chemistry",
	}
	nextID++
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// call the get handler function
		getTeachersHandler(w, r)
		w.Write([]byte("Hello GET Method on Teachers Route"))
	case http.MethodPost:
		addTeacherHandler(w, r)
		w.Write([]byte("Hello POST Method on Teachers Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Teachers Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Teachers Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Teachers Route"))
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	Path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(Path, "/")

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teachersList := make([]models.Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teachersList = append(teachersList, teacher)
			}
		}

		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teachersList),
			Data:   teachersList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		// Handle path parameters
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		teacher, exist := teachers[id]
		if !exist {
			http.Error(w, "Teacher not found", http.StatusNotFound)
		}
		json.NewEncoder(w).Encode(teacher)
	}
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}

	json.NewEncoder(w).Encode(response)
}
