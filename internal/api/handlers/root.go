package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root route"))
	fmt.Println("Hello root route")
}
