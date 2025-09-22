package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mw "academy-app-system/internal/api/middlewares"
	"academy-app-system/internal/api/router"
	"academy-app-system/pkg/utils"

	"github.com/joho/godotenv"
)

//go:embed .env
var envFile embed.FS

func loadEnvFromEmbeddedFile() {
	// Read the embedded .env file
	content, err := envFile.ReadFile(".env")
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	// Create a temp file to load the env vars
	tempfile, err := os.CreateTemp("", ".env")
	if err != nil {
		log.Fatalf("Error creating temp .env file: %v", err)
	}
	defer os.Remove(tempfile.Name())

	// Write content of the embedded .env file to the time file
	_, err = tempfile.Write(content)
	if err != nil {
		log.Fatalf("Error writing to temp .env file: %v", err)
	}

	err = tempfile.Close()
	if err != nil {
		log.Fatalf("Error closing temp file: %v", err)
	}

	// Load env vars from the temp file
	err = godotenv.Load(tempfile.Name())
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {

	// load environmnet variables from the embedded .env file
	loadEnvFromEmbeddedFile()

	fmt.Println("Environment variable CERT_FILE:", os.Getenv("CERT_FILE"))

	port := os.Getenv("API_PORT")

	cert := os.Getenv("CERT_FILE")
	key := os.Getenv("KEY_FILE")

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS10,
	}

	rl := mw.NewRateLimiter(5, time.Minute)
	HPPOptions := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		WhiteList:                   []string{"sortBy", "sortOrder", "name", "age", "class", "country"},
	}

	router := router.MainRouter()
	jwtMiddleware := mw.MiddlewaresExcludePaths(mw.JWTMiddleware, "/execs/login", "/execs/forgotpassword", "/execs/resetpassword/reset")
	secureMux := utils.ApplyMiddlewares(router, mw.SecurityHeader, mw.Compression, mw.Hpp(HPPOptions), mw.XSSMiddlware, jwtMiddleware, mw.ResponseTimeMiddleware, rl.RLMiddleware, mw.Cors)

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
