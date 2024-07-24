package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
    "strconv"
//    "log"

    "github.com/gorilla/mux"
)

// Data structure for the greeting template
//type HostData struct {
//    Name string
//}
/// Database initialization function

func isIntPrime(num int) bool {
    if num == 0 || num == 1 {
        return false
    }

    for i:=2; i<num; i++ {
        if num % i == 0 {
             return false
            }
    }
    return true
}

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

    // Define a handler function for the home page
    isPrimeHandler := func(w http.ResponseWriter, r *http.Request) {
        p := ""
        // Check if the request is a POST request
        if r.Method == http.MethodPost {
            // Parse the form data
            err := r.ParseForm()
            if err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
            }
            num := strings.TrimSpace(r.Form.Get("check"))
            intNum, err := strconv.Atoi(num)
            if err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
            }
            check := isIntPrime(intNum)
            if check == false {
                p = "the numer is NOT prime"
            } else {
                p = "the number is prime"
            }
        }

//        if r.Method == http.MethodGet {
//            p = ""
//        }
        // Parse the HTML template
        tmpl, err := template.ParseFiles("./web/templates/isprime.html")
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        // Execute the template
        err = tmpl.Execute(w, p)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
    }
    // Register the handlers for the home, greet, and login pages
    router.HandleFunc("/", homeHandler).Methods(http.MethodGet)
    router.HandleFunc("/isprime", isPrimeHandler).Methods(http.MethodGet, http.MethodPost)

    // Serve static files from the "static" directory
    //router.PathPrefix("./web/static/").Handler(http.StripPrefix("./web/static/", http.FileServer(http.Dir("static"))))

    // Start the HTTP server on port 8080 using the router
    fmt.Println("Server is listening on :8080...")
    http.ListenAndServe(":8080", router)
}
