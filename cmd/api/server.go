package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	mw "academy-app-system/internal/api/middlewares"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root route"))
	fmt.Println("Hello root route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")

		// fmt.Println("id is:", userID)

		// fmt.Println("Query params:", r.URL.Query())
		queryParams := r.URL.Query()
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortorder")

		if sortorder == "" {
			sortorder = "DESC"
		}

		fmt.Printf("Sortby: %v, Key: %v, Sortorder:%v\n", sortby, key, sortorder)
		w.Write([]byte(fmt.Sprintf("Hello GET Method on Teachers Route, the userID is %s", userID)))
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
}

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// create custom server
	server := &http.Server{
		Addr: port,
		// Handler:   mux,
		// Handler:   mw.Cors(mux),
		Handler:   mw.SecurityHeader(mw.Cors(mux)),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is listening on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}

}
