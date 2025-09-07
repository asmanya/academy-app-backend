package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello root route")
	w.Write([]byte("Hello root route"))
	fmt.Println("Hello root route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// teachers/{id}
		// teachers/9
		// teachers/?key=value&query=value2&sortby=email&sortorder=ASC
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")

		fmt.Println("id is:", userID)

		fmt.Println("Query params:", r.URL.Query())
		queryParams := r.URL.Query()
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortorder")

		if sortorder == "" {
			sortorder = "DESC"
		}

		fmt.Printf("Sortby: %v, Key: %v, Sortorder:%v\n", sortby, key, sortorder)


		w.Write([]byte("Hello GET Method on Teachers Route"))
		// fmt.Println("Hello GET Method on Teachers Route")
	case http.MethodPost:
		
		w.Write([]byte("Hello POST Method on Teachers Route"))
		fmt.Println("Hello POST Method on Teachers Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Teachers Route"))
		fmt.Println("Hello PUT Method on Teachers Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Teachers Route"))
		fmt.Println("Hello PATCH Method on Teachers Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Teachers Route"))
		fmt.Println("Hello DELETE Method on Teachers Route")
	}

	w.Write([]byte("Hello Teachers route"))
	fmt.Println("Hello Teachers route")
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method Students Route"))
		fmt.Println("Hello GET Method Students Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST Method Students Route"))
		fmt.Println("Hello POST Method Students Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method Students Route"))
		fmt.Println("Hello PUT Method Students Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method Students Route"))
		fmt.Println("Hello PATCH Method Students Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method Students Route"))
		fmt.Println("Hello DELETE Method Students Route")
	}

	w.Write([]byte("Hello Students route"))
	fmt.Println("Hello Students route")
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method execs Route"))
		fmt.Println("Hello GET Method execs Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST Method execs Route"))
		fmt.Println("Hello POST Method execs Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method execs Route"))
		fmt.Println("Hello PUT Method execs Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method execs Route"))
		fmt.Println("Hello PATCH Method execs Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method execs Route"))
		fmt.Println("Hello DELETE Method execs Route")
	}

	w.Write([]byte("Hello execs route"))
	fmt.Println("Hello execs route")
}

func main() {

	port := ":3000"

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/teachers/", teachersHandler)
	http.HandleFunc("/students/", studentsHandler)
	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is listening on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}

}

// CRUD Operations --> CRUD stands for Create, Read, Update, Delete
// HTTP methods are used to perform actions on resources in database --> POST, GET, PUT, DELETE, PATCH
// these methods corresponds to the CRUD operations