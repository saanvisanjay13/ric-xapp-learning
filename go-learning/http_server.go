package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type SecurityReport struct {
	Platform string `json: "platform"`
	Score int `json: "score"`
	Status string `json: "status"`
}

func homeHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "RIC Security Checker xApp -Running!")
}

func healthHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Status: Healthy")
}

func securityHandler(w http.ResponseWriter, r * http.Request) {
	report := SecurityReport{
		Platform: "RIC-Platform-Test",
		Score: 75,
		Status: "MOSTLY SECURE",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func  platformsHandler(w http.ResponseWriter, r * http.Request) {
	platforms := []SecurityReport{
	{Platform: "RIC-Platform-A", Score: 100, Status: "FULLY SECURE"},
	{Platform: "RIC-Platform-B", Score: 0, Status: "VULNERABLE"},
	{Platform: "RIC-Platfrom-C", Score: 75, Status: "MOSTLY SECURE"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(platforms)
}
func  versionHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "xApp Version: 1.0")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/security", securityHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/platforms" , platformsHandler)
	fmt.Println("xApp server running on port 9090...")
	http.ListenAndServe(":9090" , nil)
}









