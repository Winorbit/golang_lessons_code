package main

import ("fmt"
        str "strings")


func main () {
my_word := "Ground"
sentence :=  "In a hole in the ground there lived a hobbit."

//result := str.Contains(sentence, my_word )
//fmt.Println(result)

//result := str.Contains(str.ToLower(sentence), str.ToLower(my_word))
//fmt.Println(result)

result := str.Contains(str.ToUpper(sentence), str.ToUpper(my_word))
fmt.Println(result)

var sentence1 string; sentence1 = "The sky the port was the color of television, tuned to a dead channel."
var sentence2 string = "The man in black fled across the desert, and the gunslinger followed."
var sentence3 string
sentence3 = "We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore"
sentence4 := "In a hole in the ground there lived a hobbit."

/*
var (sentence1 string = "The sky the port was the color of television, tuned to a dead channel."
     sentence2 = "The man in black fled across the desert, and the gunslinger followed."
     sentence3 ="We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore"
     sentence4 = "In a hole in the ground there lived a hobbit."
 )
*/
result1 := str.Contains(str.ToUpper(sentence1), str.ToUpper(my_word))
fmt.Println(result1)

result2 := str.Contains(str.ToUpper(sentence2), str.ToUpper(my_word))
fmt.Println(result2)

result3 := str.Contains(str.ToUpper(sentence3), str.ToUpper(my_word))
fmt.Println(result3)

result4 := str.Contains(str.ToUpper(sentence4), str.ToUpper(my_word))
fmt.Println(result4)

}