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
	fmt.Println(r.Method, r.URL)
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	return
}

func returnAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	// tasks := Tasks{
	// 	Task{Title: "Task 1"},
	// 	Task{Title: "Task 2"},
	// }

	fmt.Println("Endpoint hit: returnAllTasks")

	json.NewEncoder(w).Encode(tasksSlice)
	return
}

func addTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: addTask")

	var t Task
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

func returnTaskStatus(w http.ResponseWriter, r *http.Request) {
	queryParamters := r.URL.Query()
	fmt.Println(r.Method, r.URL)
	fmt.Println(r.URL.Host)
	fmt.Println(queryParamters, queryParamters["test"])
	fmt.Println(r)
	w.Write([]byte(fmt.Sprintf("title:%s", r.Method)))
}

func main() {
	r := chi.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/api/task/{user}/status", returnTaskStatus)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			r.Get("/allTasks", returnAllTasks)
			r.Post("/addTask", addTask)
		})
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
