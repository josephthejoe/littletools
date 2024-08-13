package handlers

import  (
    "html/template"
    "net/http"
//    "log"                                                                     
//    "github.com/josephthejoe/littletools/internal/tools"                        
//    "github.com/gorilla/mux"
)

// Define a handler function for the home page
func HomeHandler (w http.ResponseWriter, r *http.Request) {
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
