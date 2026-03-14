package main
import "fmt"

func grade(score int) string {
	if score >= 90 {
	return "A"
	} else if score >= 80 {
	return "B"
	} else if score >= 70 {
	return "C"
	} else { 
	return "F"
	}
}

func main() {
	fmt.Println(grade(95))
	fmt.Println(grade(83))
	fmt.Println(grade(55))
	fmt.Println(grade(100))
}
