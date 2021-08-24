package main

import "fmt"

var words []string = []string{"Hello", "there", "general", "Kenobi"}
// Расположение роли не игрет - это компилируемый язык. Можно объявлять функции и под main()

func main() {
  funcWithFloatArgs(12,18, 99)
  //funcWithFloatArgs("Hello", 12,18, 99)
  //fmt.Println(returnMultipleValues())
  //emptyReturn()
  //helloThere()
}


func funcWithFloatArgs(numbers ...int){
  fmt.Println(numbers)
}


/*
func funcWithFloatArgs(numbers ...int){
  for _, number := range numbers{
    fmt.Println(number)
  }
}
*/

/*
func funcWithFloatArgs(message string, numbers ...int){
  for _, number := range numbers{
    fmt.Println(number)
  }
}
*/

/*
func returnMultipleValues()(int, int){
  return 12, 13
}
*/

/*
func emptyReturn(){
  val := "Hello there!"
  fmt.Println(val)
  return 
  // зачем нужен? Прерывание выполнения функции
}
*/

/*
func helloThere() {
  for _, word := range words{
    fmt.Println(word)
    if word == "general"{
      return
    }
  }
  fmt.Println("Никогда не будет напечатано.")
}
*/