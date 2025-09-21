package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	mw "academy-app-system/internal/api/middlewares"
	"academy-app-system/internal/api/router"
	"academy-app-system/internal/repository/sqlconnect"
	"academy-app-system/pkg/utils"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	_, err = sqlconnect.ConnectDb()
	if err != nil {
		utils.ErrorHandler(err, "")
		return
	}

	port := os.Getenv("API_PORT")

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	router := router.MainRouter()

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
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
