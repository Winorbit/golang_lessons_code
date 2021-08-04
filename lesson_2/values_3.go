package main

import "fmt"

func main() {
    /*
    Указатели это объекты, значением которых 
    служат адреса других объектов.
    То есть, если у переменной значение это, например, 45
    то у указателя - адрес ячейки памяти этого значения.
    */
    
    //var username string
    //fmt.Println(username)

    //var username *string
    //fmt.Println(username)

    username := "Alan Wake" // Определяем переменную.
    fmt.Println(username)
    fmt.Println(&username)  // Адрес этой переменной.

    var variable *string    /* Определяем указатель 
                               * говорит, что variable получит не значение
                               а именно адрес ячейки.
                            */

    variable = username    /* Вот это попытка присвоить указателю ЗНАЧЕНИЕ, 
                           так не сработает.
                           */ 
    
    variable = &username   // А вот тут мы уже присваиваем именно ссылку на ячейку памяти.

    fmt.Println(variable)
    fmt.Println(*variable)
    fmt.Println(&variable)
    
  }
