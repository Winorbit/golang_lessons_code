package main

import "fmt"

func sayHello(name string){fmt.Printf("Hello, %s \n", name)}

var message = func(name string) {fmt.Printf("Hello, %s \n", name)}

func main() {
  message("Egor")
  sayHello("Vasil")
}

