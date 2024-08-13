package handlers

import (
//    "html/template"                                                           
    "net/http"
//    "strings"                                                                 
//    "strconv"                                                                 
    "encoding/json"
//    "log"                                                                     
    "github.com/josephthejoe/littletools/internal/tools"

    "github.com/gorilla/mux"

)

func IpDataAPIHandler(w http.ResponseWriter, r *http.Request) {
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
