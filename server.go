package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	ratSelector := r.Form.Get("ratSelector")
	id := r.Form.Get("id")
	tac := r.Form.Get("tac")
	mnc := r.Form.Get("mnc")
	mcc := r.Form.Get("mcc")
	n2Information := r.Form.Get("n2Information")
	repetitionPeriod := r.Form.Get("repetitionPeriod")
	numberOfBroadcastsRequested := r.Form.Get("numberOfBroadcastsRequested")
	warningMessageContents := r.Form.Get("warningMessageContents")

	flag.Parse()
	m := make(map[string]string)
	m["ratSelector"] = ratSelector
	m["id"] = id
	m["tac"] = tac
	m["mnc"] = mnc
	m["mcc"] = mcc
	m["n2Information"] = n2Information
	m["repetitionPeriod"] = repetitionPeriod
	m["numberOfBroadcastsRequested"] = numberOfBroadcastsRequested
	m["warningMessageContents"] = warningMessageContents
	transfer(m)
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
