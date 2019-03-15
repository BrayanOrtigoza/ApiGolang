package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"API/controler"
)

func InitRouting() {



	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/taskwork", controler.NewPeoplePost).Methods("POST")

	server := &http.Server{
		Addr:            ":4000",
		Handler:         r,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		MaxHeaderBytes:  1 << 20,
	}

	log.Println("Listening http://localhost:4000 ...")
	server.ListenAndServe()
}