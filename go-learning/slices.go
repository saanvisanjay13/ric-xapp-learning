package main
import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6)
	numbers = append(numbers, 7)
	cities := []string{"Mumbai", "Bengaluru", "Delhi", "Chennai", "Kolkata"}
	for index ,city := range cities {
		fmt.Println(index, "->", city)
		}

	fmt.Println(numbers)
	fmt.Println("First:", numbers[0])
	fmt.Println("Last:", numbers[len(numbers)-1])
	fmt.Println("Length:", len(numbers))
}
