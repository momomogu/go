package main

import "fmt"

func main() {
	for {
		var command string
		fmt.Print("1.hello world other.error")
		fmt.Scan(&command)
		switch command {
		case "1":
			fmt.Println("hello world!")
		default:
			fmt.Println("error!")
		}
	}
}
