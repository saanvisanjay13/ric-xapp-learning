package main
import "fmt"
func main() {
	// Way 1 -explicit type
	var name string ="Saanvi"
	var age int= 21
	var score float64 = 9.5
	var isStudent bool = true

	// Way 2 - Go figures out the type (use this more)
	city := "Bangalore"
	year := 2025
	college := "BMSIT"
	cgpa := 8.63

	fmt.Println(name, age, score, isStudent)
	fmt.Println(city, year, college , cgpa)
}
