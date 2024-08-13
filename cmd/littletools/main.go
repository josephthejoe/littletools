package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
    "strconv"
    "encoding/json"
//    "log"
    "github.com/josephthejoe/littletools/internal/tools"
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

    // Define a handler function for the home page
//    homeHandler := func(w http.ResponseWriter, r *http.Request) {
//        // Parse the HTML template
//        tmpl, err := template.ParseFiles("./web/templates/index.html")
//        if err != nil {
//            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//            return
//        }
//
//        // Execute the template
//        err = tmpl.Execute(w, nil)
//        if err != nil {
//            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//            return
//        }
//    }

    isPrimeAPIHandler := func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]
        num, err := strconv.Atoi(id)
        if err != nil {
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        check := tools.IsIntPrime(num)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]bool{"check": check})
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
            id := strings.TrimSpace(r.Form.Get("check"))
            num, err := strconv.Atoi(id)
            if err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
            }
            check := tools.IsIntPrime(num)
            if check == false {
                p = "the numer is NOT prime"
            } else {
                p = "the number is prime"
            }
        }

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

    ipDataHandler := func(w http.ResponseWriter, r *http.Request) {
        // Check if the request is a POST request
        if r.Method == http.MethodPost {
            // Parse the form data
            err := r.ParseForm()
            if err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
            }

            ip := strings.TrimSpace(r.Form.Get("data"))
            networkA, broadcastA, cidr, mask, wc := tools.IpData(ip)
            net := NetworkData {
                Network: networkA,
                Broadcast: broadcastA,
                CIDR: cidr,
                Netmask: mask,
                Wildcard: wc,
            }
        tmpl, err := template.ParseFiles("./web/templates/ipdata.html")
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        // Execute the template
        err = tmpl.Execute(w, net)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
        }


        if r.Method == http.MethodGet {

        tmpl, err := template.ParseFiles("./web/templates/ipdata.html")
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
    }
    ipDataAPIHandler := func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        ip := vars["id"]
        networkA, broadcastA, cidr, mask, wc := tools.IpData(ip)
        net := NetworkData {
            Network: networkA,
            Broadcast: broadcastA,
            CIDR: cidr,
            Netmask: mask,
            Wildcard: wc,
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(net)
    }
    // Register the handlers for the home, greet, and login pages
    router.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)
    router.HandleFunc("/isprime", isPrimeHandler).Methods(http.MethodGet, http.MethodPost)
    router.HandleFunc("/api/isprime/{id}", isPrimeAPIHandler).Methods(http.MethodGet, http.MethodPost)
    router.HandleFunc("/ipdata", ipDataHandler).Methods(http.MethodGet, http.MethodPost)
    router.HandleFunc("/api/ipdata/{id}", ipDataAPIHandler).Methods(http.MethodGet, http.MethodPost)

    // Serve static files from the "static" directory
    //router.PathPrefix("./web/static/").Handler(http.StripPrefix("./web/static/", http.FileServer(http.Dir("static"))))

    // Start the HTTP server on port 8080 using the router
    fmt.Println("Server is listening on :8080...")
    http.ListenAndServe(":8080", router)
}
