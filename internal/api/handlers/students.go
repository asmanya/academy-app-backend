package handlers

import "net/http"

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
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
