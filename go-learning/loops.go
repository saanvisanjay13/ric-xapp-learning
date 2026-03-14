package main
import "fmt"

func main() {
	//basic loop
	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i) }
	//range loop through a list
	fruits  := []string{" apple", "banana", "mango"}
	for index , fruit  := range fruits {
	fmt.Println(index, "->" , fruit)
	}

	//while style loop
	x := 0		
	for x < 5 {
		fmt.Println ("x is:", x)
		x++
		}
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			}
			}
}
