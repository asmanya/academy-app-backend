package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	mw "academy-app-system/internal/api/middlewares"
	"academy-app-system/internal/api/router"
)

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	router := router.Router()

	// rl := mw.NewRateLimiter(5, time.Minute)
	// HPPOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	WhiteList:                   []string{"sortBy", "sortOrder", "name", "age", "class", "country"},
	// }

	// secureMux := mw.Cors(rl.RLMiddleware(mw.ResponseTimeMiddleware(mw.SecurityHeader(mw.Compression(mw.Hpp(HPPOptions)(mux))))))

	// secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(HPPOptions), mw.Compression, mw.SecurityHeader, mw.ResponseTimeMiddleware, rl.RLMiddleware, mw.Cors)

	secureMux := mw.SecurityHeader(router)

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
