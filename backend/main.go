package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pborges/tilttest/backend/api"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	db, err := sqlx.Connect("mysql", "root:dbpass@(db-service.default.svc.cluster.local:3306)/classicmodels")
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	corsHandler := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Annand")
	})

	r.HandleFunc("/customers.json", func(w http.ResponseWriter, r *http.Request) {
		customers, err := api.GetCustomers(db)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err == nil {
			json.NewEncoder(w).Encode(customers)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Serving on port 8000")
	err = http.ListenAndServe(":8000", loggingMiddleware(corsHandler))
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}
