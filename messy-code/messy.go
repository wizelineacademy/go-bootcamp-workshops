package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Todo struct (Model)
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Status    string `json:"status"`
	IsDeleted bool   `json:"isDeleted"`
}

var todos []Todo

// GET /todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// GET /todos/{id}
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

func main() {
	router := mux.NewRouter()

	// TODO get data from DB or CSV
	// todos = append(todos, Todo{ID: 10, Task: "Wash dishes", Status: "pending", IsDeleted: false})
	// todos = append(todos, Todo{ID: 20, Task: "Make report", Status: "pending", IsDeleted: false})
	// GET data from CSV
	var filename string = "./data/todos.csv"
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Unable to open CSV file!", err)
	}
	fmt.Println("Loading records from ", filename)
	r := csv.NewReader(csvfile)
	numOfRecords := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println("CSV reading done.")
			break
		}
		if err != nil {
			switch t := err.(type) {
			default:
				log.Fatalln("When reading CSV", err)
			case *csv.ParseError:
				fmt.Println("Ignoring record #", numOfRecords, t)
				continue
			}
		}
		numOfRecords++
		fmt.Println(record[0], record[1], record[2], record[3])
		// skip headers
		if numOfRecords != 1 {
			id, idErr := strconv.Atoi(strings.TrimSpace(record[0]))
			isDeleted, delErr := strconv.ParseBool(strings.TrimSpace(record[3]))
			if idErr == nil && delErr == nil {
				var task = strings.TrimSpace(record[1])
				var status = strings.TrimSpace(record[2])
				todos = append(todos, Todo{ID: id, Task: task, Status: status, IsDeleted: isDeleted})
			} else {
				fmt.Println("Ignoring record #", numOfRecords, idErr, delErr)
			}
		}
	}
	csvfile.Close()

	// Routes
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")

	// Start server
	fmt.Println("Starting server at port [3000].")
	log.Fatal(http.ListenAndServe(":3000", router))
}
