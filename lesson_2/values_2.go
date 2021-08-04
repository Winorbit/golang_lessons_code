package main

import "fmt"

// 5 вариантов задания переменной
/*
var name string = "Alan Wake"        //способ 1

var name = "Alan Wake"               //способ 2

var name string; name = "Alan Wake"  //способ 3
 
var name = string("Alan Wake")       //способ 4

name := "Alan Wake"                  //способ 5
*/


/*
Пример плохих комментариев. 
И так понятно, что это 5 способов. Подписывать каждый - бесполезно.
*/
/*
"Alan Wake"    //username
98.02          //lattitude 
101.67         //longittude
false          //subscriber

var username string = "Alan Wake" 
var lattitude bool; lattitude = 98.02
var longittude = 101.67
var subscriber = bool(false)

var (username string = "Alan Wake"
     lattitude = 98.02
     longittude = 101.67
     subscriber = false)

*/

func main() {
     /*
     При объявлении без инициализации, т.е ты здаешь переменную,
     но она изначально должна быть пустой, стоит использовать такой стиль:
     */
     var username string
     username = "Alan Wake"

     /*
     При объявлении с инициализацией, т.е ты сразу задаешь значение,
     стоит использовать такой стиль:
     */
     username := "Alan Wake"
     fmt.Println(username)

     var empty_number int     
     fmt.Println(empty_number)

     var empty_string string  
     fmt.Println(empty_string)

     var empty_bool bool       
     fmt.Println(empty_bool)
  }
