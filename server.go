package main

import (
	"login/router"
	"net/http"
	"time"
	"log"
	"login/common"
	"fmt"
)

func main() {
	fmt.Println("Server Started ..........")

	// if changing port change in homepage to serve the ui also(optional)
	router := router.NewRouter() // create routes

	router.Methods("GET", "POST", "DELETE")

	serv := &http.Server{
		Handler	: router,
		Addr	: common.ServerAddress,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout	: 15 * time.Second,
	}
	log.Fatal(serv.ListenAndServe())

	defer fmt.Println("Server Closed !!! Please restart server...... ")
}