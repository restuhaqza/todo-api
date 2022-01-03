package main

import (
	"devcode_2nd/database"
	"devcode_2nd/helper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	response := helper.FormatResponse("Success", "Welcome to TODO API", nil)
	json, _ := json.Marshal(response)
	fmt.Fprintf(w, "%s", json)
}

func main() {

	dbConfig := database.DBConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Driver:   "mysql",
		DB:       os.Getenv("MYQL_DBNAME"),
	}

	database.Connect(dbConfig)

	http.HandleFunc("/", rootHandler)

	log.Fatalln(http.ListenAndServe(":8080", nil))
	log.Println("Server listen on port 8080")
}
