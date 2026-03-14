package main
import "fmt"

func isKnownPort(port int) string {
	portDescriptions := map[int]string{
	22: "SSH",
	80: "HTTP",
	443: "HTTPS",
	8080: "App Server",
	3000: "React Dev Server",
}
if desc, exists := portDescriptions[port]; exists {
	return "Known: " + desc
}
return "Unknown port"
}

func main() {
	portDescriptions := map[int]string{
	22: "SSH",
	80: "HTTP",
	443: "HTTPS",
	8080: "App Server",
	3000: "React Dev Server",
	5432: "PostgreSQL Database",
	27017: "MongoDB",
	6379: "Redis Cache",
	9090: "Prometheus",
	
}

fmt.Println("Port 22:", portDescriptions[22])
portDescriptions[3000] = "Dev Server"

for port, description := range portDescriptions {
	fmt.Println("Port", port, "=", description)
}

if desc, exists := portDescriptions[9999]; exists {
	fmt.Println("Found:", desc)
} else {
	fmt.Println("Port 9999 not found")
}
fmt.Println(isKnownPort(80))
fmt.Println(isKnownPort(9999))
}
