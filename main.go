package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// @title User API documentation
// @version 1.0.0
// @host localhost:5000
// @BasePath /

func main() {
	// get the port
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}

	// GET /db
	http.HandleFunc("/db", DbConnect)

	// GET /
	http.HandleFunc("/", HomePage)

	// start the server
	log.Printf("Listening on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}

}

// HomePage ... Get home page
// @Summary Get Home Page
// @Description get home Page
// @Tags
// @Success 200
// @Failure 404
// @Router / [get]
func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello Mathmatters!")

}

// DbConnect ... Get db
// @Summary Get db
// @Description get db
// @Tags
// @Success 200
// @Failure 404
// @Router /db [get]
func DbConnect(w http.ResponseWriter, r *http.Request) {

	// retrieve the url
	dbURL := os.Getenv("DATABASE_URL")

	// connect to the db
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		fmt.Fprintln(w, err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "Hello db!")

	if db != nil {
		defer db.Close()
	}
}

func getPort() (string, error) {
	// the PORT is supplied by Heroku
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
