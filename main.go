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

	// allow browsers to send information the server, this will help resolve errors during testing with a browser
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness) // handler will only fire on 'GET' request, otherwise return a 405
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

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
	
}