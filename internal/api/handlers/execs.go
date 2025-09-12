package handlers

import (
	"fmt"
	"net/http"
)

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
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
