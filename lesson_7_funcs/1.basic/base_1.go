package main

import "fmt"

/*
По умолчанию каждая программа на языке Go должна содержать как минимум одну функцию 
- функцию main, которая является входной точкой в приложение:


func имя_функции (список_параметров) (типы_возвращаемых_значений){
    выполняемые_операторы
}
*/

func helloWorld(){ fmt.Println("hello World") }
/*
func helloWorld(){ 
	fmt.Println("hello World") 
}
*/

//helloWorld()

/*
func printArgument(arg string){
	fmt.Println(arg)
}
*/

// Внутри функции своя область видимости. То, что объявлено внутри не видно снаружи.
/*
func printArgument(arg string){
	var innerValue string = "Visible only inside function"
	fmt.Println(innerValue)
}
*/
// innerValue видим только внути функции. Снаружи его позвать нельзя, словим ошибку.
// fmt.Println(innerValue)


// Агументов может быть любое число, и они могут быть любого типа
/*
func printArguments(strArg string, intArg int, mapArg map[int]string){
	fmt.Println(strArg)
	fmt.Println(mapArg)
	fmt.Println(intArg)}
*/

func main() {
	/*
	okCodes := map[int]string{200: "ok",
		   	   				  201: "created",
		   	   				  202: "accepted",
		   	   				  203: "not authorized",
		   	   				  204: "no content",}
	*/
	helloWorld()
	
	//printArgument("Hello")
	//printArguments("Hello", 88, okCodes)

	/*
	func insideMain(){
		fmt.Println("Функции нельзя объявлять и внутри других функций.")
	}
	insideMain()
	*/
	
}
