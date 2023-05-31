package main

import (
	"log"
	"os/exec"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("go", "run", "NonUEN2MessageTransferRequest.go")
	output, err := cmd.Output()
	
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Command output:\n%s", output)
	w.Write(output)
}
