package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
    "strconv"
    "encoding/json"
//    "log"

    "github.com/gorilla/mux"

//    "github.com/josephthejoe/littletools/internal/handlers"
    "github.com/josephthejoe/littletools/internal/tools"
)

type NetworkData struct {
    Network     string `json:"network"`
    Broadcast   string `json:"broadcast"`
    CIDR        string `json:"cidr"`
    Netmask     string `json:"network"`
    Wildcard    string `json:"wildcard"`
}

func ipData(input string) (string, string, string, string, string){
netmasks := map[string]string{
    "32": "255.255.255.255",
    "31": "255.255.255.254",
    "30": "255.255.255.252",
    "29": "255.255.255.248",
    "28": "255.255.255.240",
    "27": "255.255.255.224",
    "26": "255.255.255.192",
    "25": "255.255.255.128",
    "24": "255.255.255.0",
    "23": "255.255.254.0",
    "22": "255.255.252.0",
    "21": "255.255.248.0",
    "20": "255.255.240.0",
    "19": "255.255.224.0",
    "18": "255.255.192.0",
    "17": "255.255.128.0",
    "16": "255.255.0.0",
    "15": "255.254.0.0",
    "14": "255.252.0.0",
    "13": "255.248.0.0",
    "12": "255.240.0.0",
    "11": "255.224.0.0",
    "10": "255.192.0.0",
    "9":  "255.128.0.0",
    "8":  "255.0.0.0",
    "7":  "254.0.0.0",
    "6":  "252.0.0.0",
    "5":  "248.0.0.0",
    "4":  "240.0.0.0",
    "3":  "224.0.0.0",
    "2":  "192.0.0.0",
    "1":  "128.0.0.0",
    "0":  "0.0.0.0",
}

wildmasks := map[string]string{
    "32": "0.0.0.0",
    "31": "0.0.0.1",
    "30": "0.0.0.3",
    "29": "0.0.0.7",
    "28": "0.0.0.15",
    "27": "0.0.0.31",
    "26": "0.0.0.63",
    "25": "0.0.0.127",
    "24": "0.0.0.255",
    "23": "0.0.1.255",
    "22": "0.0.3.255",
    "21": "0.0.7.255",
    "20": "0.0.15.255",
    "19": "0.0.31.255",
    "18": "0.0.63.255",
    "17": "0.0.127.255",
    "16": "0.0.255.255",
    "15": "0.1.255.255",
    "14": "0.3.255.255",
    "13": "0.7.255.255",
    "12": "0.15.255.255",
    "11": "0.31.255.255",
    "10": "0.63.255.255",
    "9":  "0.127.255.255",
    "8":  "0.255.255.255",
    "7":  "1.255.255.255",
    "6":  "3.255.255.255",
    "5":  "7.255.255.255",
    "4":  "15.255.255.255",
    "3":  "31.255.255.255",
    "2":  "63.255.255.255",
    "1":  "127.255.255.255",
    "0":  "255.255.255.255",
}

    //input := "172.16.4.80/24"
    slash := strings.Split(input, "-")
    cidr := slash[1]
    mask := strings.Split(netmasks[cidr], ".")
    wcMask := strings.Split(wildmasks[cidr], ".")
    ip := strings.Split(slash[0], ".")

    oct1, _ := strconv.Atoi(ip[0])
    oct2, _ := strconv.Atoi(ip[1])
    oct3, _ := strconv.Atoi(ip[2])
    oct4, _ := strconv.Atoi(ip[3])
    mask1, _ := strconv.Atoi(mask[0])
    mask2, _ := strconv.Atoi(mask[1])
    mask3, _ := strconv.Atoi(mask[2])
    mask4, _ := strconv.Atoi(mask[3])
    wMask1, _ := strconv.Atoi(wcMask[0])
    wMask2, _ := strconv.Atoi(wcMask[1])
    wMask3, _ := strconv.Atoi(wcMask[2])
    wMask4, _ := strconv.Atoi(wcMask[3])

    net1 := mask1 & oct1
    net2 := mask2 & oct2
    net3 := mask3 & oct3
    net4 := mask4 & oct4

    broad1 := wMask1 | oct1
    broad2 := wMask2 | oct2
    broad3 := wMask3 | oct3
    broad4 := wMask4 | oct4

    networkA := strconv.Itoa(net1) + "." + strconv.Itoa(net2) + "." + strconv.Itoa(net3) + "." + strconv.Itoa(net4)
    broadcastA := strconv.Itoa(broad1) + "." + strconv.Itoa(broad2) + "." + strconv.Itoa(broad3) + "." + strconv.Itoa(broad4)
    return networkA, broadcastA, cidr, netmasks[cidr], wildmasks[cidr]
}

//func isIntPrime(num int) bool {
//    if num == 0 || num == 1 {
//        return false
//    }
//
//    for i:=2; i<num; i++ {
//        if num % i == 0 {
//             return false
//            }
//    }
//    return true
//}

//func ipCalc(ipcidr string) {
//    
//}

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

    isPrimeAPIHandler := func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]
        num, err := strconv.Atoi(id)
        if err != nil {
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        check := tools.isIntPrime(num)
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
            check := tools.isIntPrime(num)
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
            networkA, broadcastA, cidr, mask, wc := ipData(ip)
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
        networkA, broadcastA, cidr, mask, wc := ipData(ip)
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
    router.HandleFunc("/", homeHandler).Methods(http.MethodGet)
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
