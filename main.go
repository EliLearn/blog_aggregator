package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // load enviornmental variables
	fmt.Println ("\n Time for learning") // necessary encouraging message


	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	//fmt.Println("Port:", portString)

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}
	
	log.Printf("Server starting on Port: %v", portString)
	//return an error if problem with handling http request
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	router.Use(cors.Handler(cors.Options{}))
	
	
}