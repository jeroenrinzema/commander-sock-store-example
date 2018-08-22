package main

import (
	"net/http"
	"os"
	"time"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/query/common"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/query/controllers"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/query/rest"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	common.OpenDatabase()
	router := common.OpenWebHub()

	router.HandleFunc("/find/{id}", rest.Use(controllers.FindByID, Authentication)).Methods("GET")
	router.HandleFunc("/find/", rest.Use(controllers.FindAll, Authentication)).Methods("GET")

	server := &http.Server{
		Addr:         os.Getenv("HOST_ADDRESS"),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}

// Authentication validates if the given request is authenticated.
// If the request is not authenticated is a 401 returned.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// <- authenticate request and provide the users dataset key
		next.ServeHTTP(w, r)
	}
}
