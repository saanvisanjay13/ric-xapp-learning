package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Platform struct {
	Name string `json:"name"`
	Encrypted bool `json:"encrypted"`
	Firewall bool `json:"firewall"`
	OpenPorts int `json:"open_ports"`
	AuthEnabled bool `json:"auth_enabled"`
}

type SecurityResult struct {
	Platform string `json:"platform"`
	Score int `json:"score"`
	Status string `json:"status"`
	Details string `json:"details"`
}

func checkSecurity(p Platform) SecurityResult {
	score := 0
	details := ""
	if p.Encrypted { score += 25; details += "Encrypted OK " } else
	{ details += "No encryption " }
	 if p.Firewall { score += 25; details += "Firewall OK " } else { details += "No firewall " }
	if p.OpenPorts < 5 { score += 25; details += "Ports OK " } else { details += "Too many ports "}
	if p.AuthEnabled { score += 25; details += "Auth OK" } else { details += "No Auth" }
	status := "VULNERABLE"
	if score == 100 { status = "FULLY SECURE" } else if score >= 75 {status = "MOSTLY SECURE" } else if score >= 50 { status = "MODERATE RISK" }
	return SecurityResult{Platform: p.Name, Score: score, Status: status, Details: details}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RIC Security xApp v1.0 - Running")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	platforms := []Platform{
	{Name: "RIC-Node-1", Encrypted: true, Firewall: true,OpenPorts: 3, AuthEnabled: true},
	{Name: "RIC-Node-2", Encrypted: false, Firewall: true, OpenPorts: 8, AuthEnabled: false},
	{Name: "RIC-Node-3", Encrypted: true, Firewall: false, OpenPorts: 2, AuthEnabled: true},
}

var results []SecurityResult
for _, p := range platforms {
	results = append(results, checkSecurity(p))
}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/check", checkHandler)
	fmt.Println("RIC Security xApp starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
