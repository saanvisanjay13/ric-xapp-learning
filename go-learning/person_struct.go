package main
import "fmt"

type Person  struct{
	Name string
	Age int
	Email string
}

func main(){
	p1 := Person{
		Name: "Saanvi",
		Age: 20,
		Email: "saanvi@example.com",
}
fmt.Println("Name:", p1.Name)
fmt.Println("Age:", p1.Age)
p1.Age = 21
fmt.Println("Updated age:", p1.Age)
}
