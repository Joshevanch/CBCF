package main

import (
	"log"
	"os/exec"
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/notify", handleNotify)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("go", "run", "NonUeN2InfoSubscribe.go")
	output, err := cmd.Output()
	
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Subscribe output:\n%s", output)
	w.Write(output)

	cmd = exec.Command("go", "run", "NonUEN2MessageTransferRequest.go")
	output, err = cmd.Output()
	
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Message Transfer output:\n%s", output)
	w.Write(output)
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