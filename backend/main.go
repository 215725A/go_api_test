package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the Golang backend!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "Golang backend!",
		"status":  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS設定

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func databaseHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Database query Error:", err)
		return
	}
	defer rows.Close()

	var userData []map[string]interface{}

	for rows.Next() {
		var id int
		var name string
		var email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userData = append(userData, map[string]interface{}{
			"ID":    id,
			"Name":  name,
			"Email": email,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err = json.NewEncoder(w).Encode(userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func main() {
	connectionString := "host=db user=user dbname=users password=password sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", handler)
	http.HandleFunc("/api/data", apiHandler)
	http.HandleFunc("/api/db", func(w http.ResponseWriter, r *http.Request) {
		databaseHandler(w, r, db)
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // フロントエンドのオリジンを設定
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(http.DefaultServeMux)))
}
