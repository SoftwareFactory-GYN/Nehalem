package main

import (
	"fmt"
	"github.com/SoftwareFactory-GYN/nehalem/rest_api/db"
	"github.com/gocql/gocql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var CassandraSession *gocql.Session

func initRouter(r *mux.Router) {

	// Endpoint used to serve login page
	r.Handle("/api/login", LoginHandler).Methods("POST")

	// Endpoint for registering
	r.Handle("/api/register", RegisterHandler).Methods("POST")

	//Index endpoint
	r.Handle("/", jwtMiddleware.Handler(IndexHandler)).Methods("GET")
}

func main() {

	// Cassandra init
	db.CassandraInit()
	CassandraSession = db.GetSession()
	defer CassandraSession.Close()

	r := mux.NewRouter()

	// Add our routes
	initRouter(r)

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	fmt.Println("Serving on: http://localhost:3000")
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}
