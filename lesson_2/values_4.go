package main

import ("fmt"
        "strings")

func main(){
/*
2+3
fmt.Println(2+3)
fmt.Println(2*3)
fmt.Println(2-3)
*/

/*
someFloat := 2.0
fmt.Println(someFloat+3.7)
fmt.Println(2.5*3)
fmt.Println(33/7)
var someFloat32 = float32(33)
fmt.Println(someFloat32/float32(7))

//Нельзя использовать разные типы данных - в Go нет автоматического преобразования
fmt.Println(float32(3)/float64(77))
*/

/*
// Преобразование в bool
fmt.Println(2 == 3)
fmt.Println(2 == 2)

fmt.Println(2 > 3)
fmt.Println(2 < 3)
fmt.Println(2 <= 3)
*/

/*
myWord := "Hello"
fmt.Println(myWord == "world")
fmt.Println(myWord == "Hello")

fmt.Println("Hello" + "world")
fmt.Println("Hello "+ "world")
fmt.Println("Hello" + " " + "world")

fmt.Println(myWord * 2)
fmt.Println("Hello"+" "+"world")
*/

// String format
/*
fmt.Printf("Hello %d", 10)
fmt.Printf("Hello %d", "10")
fmt.Printf("Hello %d", "world")
fmt.Printf("Hello %s", "world")
fmt.Printf("Hello %v", "world")


res := strings.Contains("Word", "wor") 
fmt.Println(res)

// Размер буквы важен, S и s это разные символы.

res := strings.Contains("Word", "Wor") 
fmt.Println(res)

fmt.Println(strings.ToLower("Word"))
fmt.Println(strings.ToUpper("Word"))
*/

}