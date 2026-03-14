package main
import "fmt"

func add(a int, b int) int {
	return a + b
}

func mult(a int, b int) int {
	return a * b
}

func sub(a int, b int) int {
	return a - b
}

func greet(name string ) string {
	return "Hello, " + name + "!"
}

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	result  := add(5, 3)
	result1 := mult(8, 3)
	result2 := sub(8, 3)
	message  := greet("Saanvi")
	even := isEven(4)
	fmt.Println("Sum:", result)
	fmt.Println(message)
	fmt.Println("Is 4 even?", even)
	fmt.Println("Product:", result1)
	fmt.Println("Difference:", result2)
}
