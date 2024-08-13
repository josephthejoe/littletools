package handlers

import (
    "net/http"
    "strconv"
    "encoding/json"
//    "log"                                                                     
    "github.com/josephthejoe/littletools/internal/tools"
    "github.com/gorilla/mux"
)

func IsPrimeAPIHandler(w http.ResponseWriter, r *http.Request) {
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
