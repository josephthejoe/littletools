package handlers

import (
    "html/template"
    "net/http"
    "strings"
    "github.com/josephthejoe/littletools/internal/tools"

)

type NetworkData struct {
    Network     string `json:"network"`
    Broadcast   string `json:"broadcast"`
    CIDR        string `json:"cidr"`
    Netmask     string `json:"network"`
    Wildcard    string `json:"wildcard"`
}

func IpDataHandler (w http.ResponseWriter, r *http.Request) {
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
