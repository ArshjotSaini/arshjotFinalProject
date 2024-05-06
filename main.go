package main

import (
	"arshjotFinalProject/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

const IMAGEPATH = "./assets/nameofimage.png"

func main() {
	utils.Dbinit()
	utils.Mapinit(IMAGEPATH)
	// Create a new router
	router := mux.NewRouter()

	// registering the handler and routing
	router.HandleFunc("/", utils.HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/map", utils.MapHandler).Methods(http.MethodGet, http.MethodPost)

	// Serve static files from the "static" directory
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Start the HTTP server on port 9090 using the router
	fmt.Println("Server is listening on :9090...")
	http.ListenAndServe(":9090", router)
}
