package handlers

import (
    "html/template"
    "net/http"
    "strings"
    "strconv"
//    "log"
    "github.com/josephthejoe/littletools/internal/tools"
)

func IsPrimeHandler(w http.ResponseWriter, r *http.Request) {
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
