package main

import (
	"log"
	"net/http"
	"os"
	"constant"
	"db"
	"routes"
	"authorization"


	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

func main() {
	//Database connection
	database.DbConnect()
	defer database.Db.Close()
	port := ":" + os.Getenv("PORT")
	//port := ":8000"
	//Init Router
	r := mux.NewRouter()

	//User routes
	r.HandleFunc("/api/user/create", routes.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/verify", routes.VerifyUser).Methods("POST")
	r.HandleFunc("/api/user/getAnswers/{id}", routes.GetAnswers).Methods("GET")
	//Match routes
	r.Handle("/api/match/create", authorization.IsAuthorized(routes.CreateMatch, constant.UserAdmin)).Methods("POST")

	log.Fatal(http.ListenAndServe(port, r))
}