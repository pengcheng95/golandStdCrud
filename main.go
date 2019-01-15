package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	Title string `json:"Title"`
}

type Tasks []Task

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := Tasks{
		Task{Title: "Task 1"},
		Task{Title: "Task 2"},
	}

	fmt.Println("Endpoint hit: returnAllTasks")

	json.NewEncoder(w).Encode(tasks)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllTasks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
