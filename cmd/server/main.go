package main

import "fmt"

//	Run - is going to be responsible for the instantiation and start
//	of our Go Application
func Run() error {
	fmt.Println("Starting Up Our Application")
	return nil
}

func main() {
	fmt.Println("Comment API in Golang")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
