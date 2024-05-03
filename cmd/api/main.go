package main

import (
	"fmt"
	"net/http"

	"github.com/dev-by-sjb/go_api/internal/handlers"
	"github.com/go-chi/chi"          // a package for web dev
	log "github.com/sirupsen/logrus" // to log errors for Debugging
)

func main() {
	log.SetReportCaller(true) // this basically allow for the log to show the file name and line number of the error when i log from a file

	var r *chi.Mux = chi.NewRouter() // returns a router to a mux pointer
	handlers.Handler(r)

	fmt.Println("Starting server on: http://localhost:3000 ....")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Error(err)
	}

}
