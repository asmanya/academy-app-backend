package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {

	fmt.Println("Trying to connect MariaDB")

	/*
		WE NEED TO LOAD THE GODOTENV ONCE, THAT CAN BE DONE IN SERVER.GO FILE
			err := godotenv.Load()
			if err != nil {
				return nil, err
			}
	*/

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// username:password@tcp(remote ip of database)/database name
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, dbport, dbname)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		// panic(err)
		return nil, err
	}

	fmt.Println("Connected to MariaDB")
	return db, nil

}
