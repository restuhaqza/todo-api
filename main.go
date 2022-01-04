package main

import (
	"devcode_2nd/activity_group"
	"devcode_2nd/database"
	"devcode_2nd/helper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
		DB:       os.Getenv("MYSQL_DBNAME"),
	}

	db := database.Connect(dbConfig)

	activityGroupRepository := activity_group.NewRepository(db)

	activityGroupService := activity_group.NewService(activityGroupRepository)

	activityGroupHandler := activity_group.NewHandler(activityGroupService)

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/activity-groups", activityGroupHandler.GetAll).Methods("GET")
	r.HandleFunc("/activity-groups/{id}", activityGroupHandler.GetByID).Methods("GET")
	r.HandleFunc("/activity-groups", activityGroupHandler.Create).Methods("POST")
	r.HandleFunc("/activity-groups/{id}", activityGroupHandler.Delete).Methods("DELETE")
	r.HandleFunc("/activity-groups/{id}", activityGroupHandler.Update).Methods("PATCH")

	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(fmt.Sprintf("%s %s", r.Method, r.URL))
			h.ServeHTTP(w, r)
		})
	})

	http.Handle("/", r)

	log.Fatalln(http.ListenAndServe(":3000", nil))
	log.Println("Server listen on port 3000")
}
