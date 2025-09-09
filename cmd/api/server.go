package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	mw "academy-app-system/internal/api/middlewares"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root route"))
	fmt.Println("Hello root route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Teachers Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Teachers Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Teachers Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Teachers Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Teachers Route"))
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method Students Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST Method Students Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method Students Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method Students Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method Students Route"))
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method execs Route"))
	case http.MethodPost:
		fmt.Println("Query:", r.URL.Query())
		fmt.Println("name:", r.URL.Query().Get("name"))

		// Parse form data (necessary for x-www-form-urlencoded)
		err := r.ParseForm()
		if err != nil {
			return
		}

		fmt.Println("Form from POST methods:", r.Form)

		w.Write([]byte("Hello POST Method execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method execs Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method execs Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method execs Route"))
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

	rl := mw.NewRateLimiter(5, time.Minute)
	hpp := mw.HPPOptions{
		CheckQuery: true,
		CheckBody: true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		WhiteList: []string{"sortBy", "sortOrder", "name", "age", "class", "country"},
	}

	secureMux := mw.Hpp(hpp)(rl.RLMiddleware(mw.Compression(mw.ResponseTimeMiddleware(mw.SecurityHeader(mw.Cors(mux))))))

	// create custom server
	server := &http.Server{
		Addr: port,
		// Handler:   mux,
		// Handler:   mw.Cors(mux),
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is listening on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}

}