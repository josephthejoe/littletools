package tools
import (
    "strings"
    "strconv"
)

//type NetworkData struct {
//    Network     string `json:"network"`
//    Broadcast   string `json:"broadcast"`
//    CIDR        string `json:"cidr"`
//    Netmask     string `json:"network"`
//    Wildcard    string `json:"wildcard"`
//}

func IpData(input string) (string, string, string, string, string){
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

