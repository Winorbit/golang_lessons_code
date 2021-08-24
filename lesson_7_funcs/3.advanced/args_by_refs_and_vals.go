package main 

import "fmt"

var test int = 42 
var testSlice []int = []int{8,15,16,23}
var englishLevels map[int]string = map[int]string{1: "Beginer",
												                          2: "Middle"}

func changeByReference(val *int){*val = 99}

func changeByValue(val int){val = 99}

func addByReference(arg *[]int){
	*arg = append(*arg, 42)}

func addByValue(arg []int){
	arg = append(arg, 42)}

func updateMap(arg map[int]string){
	arg[3] = "Intermediate"}

func main(){
  /*
  changeByReference(&test)
  fmt.Println(test)
  */
  
  changeByValue(test)
  fmt.Println(test)

  //addByReference(testSlice)
  /*
  fmt.Println(testSlice)
  addByReference(&testSlice)
  fmt.Println(testSlice)
  */

  /*
  fmt.Println(testSlice)
  addByValue(testSlice)
  fmt.Println(testSlice)
  */

  // Да, maps передаються по ссылке
  fmt.Println(englishLevels)
  updateMap(englishLevels)
  fmt.Println(englishLevels)

}