package main
import "fmt"

type RICPlatform struct {
	Name string
	IsEncrypted bool
	HasFirewall bool
	OpenPorts int
	AuthEnabled bool
	TLSVersion string
}

func (r RICPlatform) SecurityScore() int {
	score := 0
	if r.IsEncrypted { score += 25 }
	if r.HasFirewall { score += 25 }
	if r.OpenPorts < 5 { score += 25 }
	if r.AuthEnabled { score += 25 }
	return score
}

func (r RICPlatform) SecurityStatus() string {
	score := r.SecurityScore()
	if score == 100 { return "FULLY SECURE" }
	if score >= 75 { return "MOSTLY SECURE" }
	if score >= 50 { return "MODERATE RISK" }
	return "VULNERABLE"
}

func (r RICPlatform) Report() {
	fmt.Println("=============")
	fmt.Println("Platform:", r.Name)
	fmt.Println("Score:", r.SecurityScore(), "/100")
	fmt.Println("Status:", r.SecurityStatus())
	fmt.Println("Encrypted:", r.IsEncrypted)
	fmt.Println("Firewall:", r.HasFirewall)
	fmt.Println("Open Ports:", r.OpenPorts)
	fmt.Println("Auth Enabled:", r.AuthEnabled)
	fmt.Println("======================")
	fmt.Println("TLS Version:" , r.TLSVersion)
}

func main() {
	platform1 := RICPlatform{
		Name: "RIC-Platform-A",
		IsEncrypted: true, HasFirewall: true, OpenPorts: 3, AuthEnabled: true, }
	platform2 := RICPlatform {
		Name: "RIC-Platform-B",
		IsEncrypted: false, HasFirewall: false, OpenPorts: 15, AuthEnabled: false, }
	platform3 := RICPlatform{
		Name: "RIC-Platform-C",
		IsEncrypted: true, HasFirewall: true, OpenPorts: 2, AuthEnabled: false,
		TLSVersion: "TLS 1.2",
		}
	platform4 := RICPlatform{
		Name: "RIC-Platform-D",
		IsEncrypted: false, HasFirewall: true, OpenPorts: 8, AuthEnabled: true,
		TLSVersion: "TLS 1.3",
		}		
	platform1.Report()
	platform2.Report()
	platform3.Report()
	platform4.Report()
}

