package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"time"
	"encoding/json"
	"io/ioutil"
)

type Todo struct {
	Name      string `json:"name"`
	Completed bool `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

var members = []Todo{{Name: "first"}}

var todosMap = make(map[string]Todo)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todos", TodoIndex).Methods("GET")
	router.HandleFunc("/todos/{todoId}", TodoShow).Methods("GET")
	router.HandleFunc("/todos/{todoId}", TodoPut).Methods("PUT")
	router.HandleFunc("/todos/{todoId}", TodoDelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
func TodoDelete(writer http.ResponseWriter, request *http.Request) {
	delete(todosMap, mux.Vars(request)["todoId"])
	writer.WriteHeader(http.StatusNoContent)
}
func TodoPut(writer http.ResponseWriter, request *http.Request) {

	var t Todo
	b, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(b, &t)
	members = append(members, t)
	todosMap[mux.Vars(request)["todoId"]] = t
	j, _ := json.Marshal(t)
	writer.Write(j)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todosMap)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	todo := todosMap[vars["todoId"]]
	fmt.Fprintln(w, "Todo show:", todo)
}