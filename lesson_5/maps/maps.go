package main

import "fmt"


/*
помещение | статус сигнализации
-----------------------
45  	  | true
-----------------------
01   	  | false
-----------------------
56	 	  | false
-----------------------


Таблица кодов ответов:

200  | "ok"
-----------------------
201  | "created"
-----------------------
202  | "accepted"
-----------------------
203  | "not authorized"
-----------------------
204  | "no content"

*/

func main() {
	okCodes := map[int]string{200: "ok",
		   	   				  201: "created",
		   	   				  202: "accepted",
		   	   				  203: "not authorized",
		   	   				  204: "no content",}

    // GET
    
	//fmt.Println(okCodes)
	//fmt.Println(okCodes[202])

	//code := okCodes[202] 
	//fmt.Println(code)

	// Что будет, если обратимся по несуществующему ключу? Пустота
	//fmt.Println(okCodes[999])
	

	// А вот так - ошибка
	//fmt.Println(okCodes["999"])

	// Update + Add
	
	fmt.Println(okCodes[204])
	okCodes[204] = "Hello"
	fmt.Println(okCodes[204])
	fmt.Println(okCodes)

	okCodes[208] = "already reported"
	fmt.Println(okCodes)
	
	
	// А вот так - нельзя
	//okCodes[208] = 802 

	/*
	// REMOVE
	fmt.Println(okCodes[208])
	delete(okCodes, 208)
	fmt.Println(okCodes[208])
	*/
}
