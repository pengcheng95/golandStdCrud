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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	return
}

func returnAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	tasks := Tasks{
		Task{Title: "Task 1"},
		Task{Title: "Task 2"},
	}

	fmt.Println("Endpoint hit: returnAllTasks")

	json.NewEncoder(w).Encode(tasks)
	return
}

func returnTaskStatus(w http.ResponseWriter, r *http.Request) {
	queryParamters := r.URL.Query()
	fmt.Println(r.Method, r.URL)
	fmt.Println(queryParamters, queryParamters["test"])
	fmt.Println(r)
	w.Write([]byte(fmt.Sprintf("title:%s", r.Method)))
}

func main() {
	r := chi.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/all", returnAllTasks)
	r.HandleFunc("/api/task/{user}/status", returnTaskStatus)
	log.Fatal(http.ListenAndServe(":8080", r))
}
