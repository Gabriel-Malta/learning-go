package main

import "fmt"

func Hello(nome string) string {
	return "Hello " + nome + "!"
}

func main()  {
	fmt.Println(Hello("World"))
}