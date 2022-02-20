package main

import "fmt"

const prefixoHelloEN = "Hello, "

func Hello(nome string) string {
	if nome == ""{
		nome = "World"
	}
	return prefixoHelloEN + nome + "!"
}

func main()  {
	fmt.Println(Hello("World"))
}