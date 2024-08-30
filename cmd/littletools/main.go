package main

import (
    "fmt"
    "net/http"
    "github.com/josephthejoe/littletools/internal/handlers"

    "github.com/gorilla/mux"
)

type NetworkData struct {
    Network     string `json:"network"`
    Broadcast   string `json:"broadcast"`
    CIDR        string `json:"cidr"`
    Netmask     string `json:"network"`
    Wildcard    string `json:"wildcard"`
}

func main() {
    // Create a new router
    router := mux.NewRouter()

    // Register the handlers for the home, greet, and login pages
    router.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)
    router.HandleFunc("/isprime", handlers.IsPrimeHandler).Methods(http.MethodGet, http.MethodPost)
    router.HandleFunc("/api/isprime/{id}", handlers.IsPrimeAPIHandler).Methods(http.MethodGet, http.MethodPost)
    router.HandleFunc("/ipdata", handlers.IpDataHandler).Methods(http.MethodGet, http.MethodPost)
    router.HandleFunc("/api/ipdata/{id}", handlers.IpDataAPIHandler).Methods(http.MethodGet, http.MethodPost)

    // Serve static files from the "static" directory
    //router.PathPrefix("./web/static/").Handler(http.StripPrefix("./web/static/", http.FileServer(http.Dir("static"))))

    // Start the HTTP server on port 8080 using the router
    fmt.Println("Server is listening on :8080...")
    http.ListenAndServe(":8080", router)
}
