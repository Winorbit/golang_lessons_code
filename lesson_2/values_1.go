package main

import "fmt"

func main() {

// Пользователь пиложения-шагомера:
"Alan Wake" //username
98.02  		//lattitude 
101.67 		//longittude
false  		//subscriber

// Песня из музыкального сервиса:
"Mockingbird" // song name
"Eminem"	  // author
"Encore"	  // album
4.11          // length


// Блюдо из сервиса заказов:
"Buritto" 		// name
8.16 	  		// weight
"5edb0382e0"	// ID
5               // count


fmt.Print("Alan Wake", 98.02, 101.67, false)
//
fmt.Print("Mockingbird")
fmt.Print("Eminem")
fmt.Print("Encore")
fmt.Print(4.11)
//
fmt.Println(8.16)
fmt.Println("Burito")
fmt.Println(5)
fmt.Println("5edb0382e0")
fmt.Println(true)

fmt.Println(type(true))
fmt.Println(type("5edb0382e0"))
fmt.Println(type(5))


// Преобразования

string(5)
int("5")
bool("5")
string(false)
string(78.79)

}

