package main

import (
	"GolangRestApi/apis/userApi"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/find", userApi.FindUser).Methods("GET")
	router.HandleFunc("/api/v1/user/getall", userApi.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/user/create", userApi.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update", userApi.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete", userApi.Delete).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err)
	}
}
