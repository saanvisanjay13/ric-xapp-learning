package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
	return 0, errors.New("cannot divide by zero")
}
return a/b, nil
}

func checkPort(port int) (string, error) {
	if  port < 1 || port > 65535 {
	return "", errors.New("invalid port number")
	}
	if port < 1024 {
		return "system port - needs root" , nil
	}
	return "user port -open", nil
}

func main() {
	result, err := divide(10,2)
	if err != nil {
	fmt.Println("Error:", err)
	}else {
	fmt.Println("Result:", result)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Result:", result2)
}

status, err3 := checkPort(8080)
if err3 != nil {
	fmt.Println("Error:", err3)
} else {
	fmt.Println("Port 8080:", status)
}
status4, err4 := checkPort(99999)
if err4 != nil {
	fmt.Println("Error:",err4)
} else {
	fmt.Println("Port 99999:", status4)
}

status5, err5 := checkPort(80)
if err5 != nil {
	fmt.Println("Error:", err5)
} else {
	fmt.Println("Port 80:", status5)
}
}
