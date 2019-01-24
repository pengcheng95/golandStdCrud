package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Task struct {
	Title string `json:"Title"`
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
		})
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
