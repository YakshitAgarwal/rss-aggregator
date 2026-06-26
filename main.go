package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		//specifies which websites (origins) are allowed to access the API
		AllowedOrigins: []string{"https://*", "http://*"}, 

		//the HTTP methods browsers are allowed to use when making cross-origin requests
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, 

		//any request header is allowed
		AllowedHeaders: []string{"*"}, 

		ExposedHeaders: []string{"Link"},
		
		//do not allow browsers to send credentials with cross-origin requests, 
		//should be false in case of REST APIs using JWT
		//should be true in case of web applications using cookies
		AllowCredentials: false, 

		//the browser remembers the preflight result for 300 seconds
		MaxAge: 300, 
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}