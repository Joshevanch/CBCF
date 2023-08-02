package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Define the Go struct representing the XML data
type Alert struct {
	XMLName    xml.Name  `xml:"urn:oasis:names:tc:emergency:cap:1.1 alert"`
	Identifier string    `xml:"identifier"`
	Sender     string    `xml:"sender"`
	Sent       time.Time `xml:"sent"`
	Status     string    `xml:"status"`
	MsgType    string    `xml:"msgType"`
	Scope      string    `xml:"scope"`
	Source     string    `xml:"source"`
	Info       Info      `xml:"info"`
}

type Info struct {
	Language     string    `xml:"language"`
	Category     string    `xml:"category"`
	Event        string    `xml:"event"`
	ResponseType string    `xml:"responseType"`
	Urgency      string    `xml:"urgency"`
	Severity     string    `xml:"severity"`
	Certainty    string    `xml:"certainty"`
	Expires      time.Time `xml:"expires"`
	SenderName   string    `xml:"senderName"`
	Headline     string    `xml:"headline"`
	Description  string    `xml:"description"`
	Contact      string    `xml:"contact"`
	Area         Area      `xml:"area"`
}

type Area struct {
	AreaDesc string `xml:"areaDesc"`
	Polygon  string `xml:"polygon"`
	GeoCode  string `xml:"geocode"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/notify", handleNotify)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	xmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error parsing body data", http.StatusBadRequest)
		return
	}
	taiwanTimezone, err := time.LoadLocation("Asia/Taipei")
	currentTime := time.Now().In(taiwanTimezone)
	fmt.Println("Time received data: ", currentTime.Format("2006-01-02 15:04:05"))
	var alertData Alert
	if err := xml.Unmarshal(xmlData, &alertData); err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing XML data", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("XML data received successfully"))
	data := make(map[string]string)
	serialNumberInteger, err := strconv.Atoi(alertData.Identifier[len(alertData.Identifier)-3:])
	serialNumber := int64(serialNumberInteger)
	serialNumberBits := strconv.FormatInt(int64(serialNumber), 2)
	serialNumberBits = "01" + "01" + serialNumberBits + "0000"
	serialNumber, err = strconv.ParseInt(serialNumberBits, 2, 64)
	data["serialNumber"] = fmt.Sprintf("%x", serialNumber)
	data["messageType"] = alertData.MsgType
	if alertData.Info.Language == "en-US" {
		data["dataCodingScheme"] = "01"
	}
	if alertData.Info.Language == "zh-TW" {
		data["dataCodingScheme"] = "48"
	}
	switch {
	case alertData.Info.Severity == "Extreme" && alertData.Info.Urgency == "Immediate" && alertData.Info.Certainty == "Observed":
		data["messageIdentifier"] = "1113"
	case alertData.Info.Severity == "Extreme" && alertData.Info.Urgency == "Immediate" && alertData.Info.Certainty == "Likely":
		data["messageIdentifier"] = "1114"
	case alertData.Info.Severity == "Extreme" && alertData.Info.Urgency == "Expected" && alertData.Info.Certainty == "Observed":
		data["messageIdentifier"] = "1115"
	case alertData.Info.Severity == "Extreme" && alertData.Info.Urgency == "Expected" && alertData.Info.Certainty == "Likely":
		data["messageIdentifier"] = "1116"
	case alertData.Info.Severity == "Severe" && alertData.Info.Urgency == "Immediate" && alertData.Info.Certainty == "Observed":
		data["messageIdentifier"] = "1117"
	case alertData.Info.Severity == "Severe" && alertData.Info.Urgency == "Immediate" && alertData.Info.Certainty == "Likely":
		data["messageIdentifier"] = "1118"
	case alertData.Info.Severity == "Severe" && alertData.Info.Urgency == "Expected" && alertData.Info.Certainty == "Observed":
		data["messageIdentifier"] = "1119"
	case alertData.Info.Severity == "Severe" && alertData.Info.Urgency == "Expected" && alertData.Info.Certainty == "Likely":
		data["messageIdentifier"] = "111A"
	default:
		data["messageIdentifier"] = "1112"
	}
	data["warningMessageContents"] = alertData.Sent.Format("2006-01-02 15:04:05") + alertData.Info.Headline + "\n" + alertData.Info.Description + "\n" + alertData.Info.Area.AreaDesc
	data["tac"] = alertData.Info.Area.GeoCode
	subscribe()
	transfer(data)
}

func handleNotify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	fmt.Println(string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Received the request body"))
}
