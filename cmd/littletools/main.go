package main

import (
    "fmt"
    "html/template"
    "net/http"
//    "log"

    "github.com/gorilla/mux"
)

// Data structure for the greeting template
//type HostData struct {
//    Name string
//}
/// Database initialization function

func main() {
    // Create a new router
    router := mux.NewRouter()

    // Define a handler function for the home page
    homeHandler := func(w http.ResponseWriter, r *http.Request) {
        // Parse the HTML template
        tmpl, err := template.ParseFiles("./web/templates/index.html")
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        // Execute the template
        err = tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
    }

    // Register the handlers for the home, greet, and login pages
    router.HandleFunc("/", homeHandler).Methods(http.MethodGet)
    //router.HandleFunc("/updatehost", updateHostHandler).Methods(http.MethodGet, http.MethodPost)

    // Serve static files from the "static" directory
    //router.PathPrefix("./web/static/").Handler(http.StripPrefix("./web/static/", http.FileServer(http.Dir("static"))))

    // Start the HTTP server on port 8080 using the router
    fmt.Println("Server is listening on :8080...")
    http.ListenAndServe(":8080", router)
}
