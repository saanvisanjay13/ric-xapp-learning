package main
import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

func (p Person) isAdult() bool {
	return p.Age >= 18
}

func main() {
	p := Person{Name: "Saanvi", Age: 21}
	fmt.Println(p.Greet())
	fmt.Println("Is Adult?", p.isAdult())
}

