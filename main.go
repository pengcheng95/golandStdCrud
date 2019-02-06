package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Task struct {
	Title  string `json:"Title"`
	Active bool   `json:"Active"`
}

type Tasks []Task

var tasksSlice = make(Tasks, 0)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
	return
}

func returnAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllTasks")
	fmt.Println(r.Method, r.URL)

	json.NewEncoder(w).Encode(tasksSlice)
	return
}

func addTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: addTask")

	var t Task
	// Checks if a request body was sent
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	tasksSlice = append(tasksSlice, t)
	w.Write([]byte(fmt.Sprintf("Task Added")))
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: deleteTask")

	// If no request body return error
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	// if unable to decode object into struct return error
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Loop through task slice to find task to be deleted
	deleted := false
	var deletedTask Task
	for i := 0; i < len(tasksSlice); i++ {
		fmt.Println(tasksSlice[i])
		curr := tasksSlice[i]
		if curr.Title == t.Title {
			deleted = true
			deletedTask = curr
			tasksSlice = append(tasksSlice[:i], tasksSlice[i+1:]...)
		}
	}

	if deleted {
		// Marshal changes struct to JSON
		deletedTaskJSON, err := json.Marshal(deletedTask)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write(deletedTaskJSON)
	} else {
		w.Write([]byte(fmt.Sprintf("Task Not Found")))
	}
}

func main() {
	r := chi.NewRouter()
	r.HandleFunc("/", homePage)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			r.Get("/allTasks", returnAllTasks)
			r.Put("/addTask", addTask)
			r.Post("/deleteTask", deleteTask)
		})
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
