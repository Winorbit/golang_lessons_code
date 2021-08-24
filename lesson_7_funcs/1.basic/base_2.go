package main

import "fmt"

func returnExample() int{return 42}

/*
func returnExample(message string) string{
	return message
}
*/

/*
func returnExample(message string) int{
	return 42
}
*/

/*
func funcWithReturn(lat float32, lon float32) string {
	coordMessage := fmt.Sprintf("coords now: lattitude %2f, longittude :%2f", lat, lon)
	return coordMessage
}
*/

/*
func funcWithReturn(lat, lon float32) string {
	coordMessage := fmt.Sprintf("coords now: lattitude %2f, longittude :%2f", lat, lon)
	return coordMessage
}
*/


func main() {
	
	returnExample()
	fmt.Println(returnExample())

	value := returnExample()
	fmt.Println(value)
	

	/*
	value := returnExample("Hello!")
	fmt.Println(value)
	*/

	/*
	coords := funcWithReturn(100.15, 93.81)
	fmt.Println(coords)
	*/

	/*
	text := fmt.Sprintf("Hello %s", "world")
	fmt.Println(text)
	*/

}
