package main

import "fmt"

/*
Опишем нашего юзера

-----------------------
name  	  		  | Jason Born
-----------------------
cardIsActive	  | false
-----------------------
lattitude 		  | 120.06
-----------------------
longitude 		  | 99.13




Опишем юзера прилаги

-----------------------
name  	  		  | Foma Kiniaev
-----------------------
englishLevel	  | Beginer



ITEM
-----------------------
id  	  | 45 <int>
-----------------------
name	  | pizza <string>
-----------------------
price	  | 15.5


ORDER
-----------------------
id  	    | ca82a6dff817e <string>
-----------------------
adress	    | "Sunset Boulevard, Str, 19/50" <string>
-----------------------
lat	  	    | 101.87
-----------------------
lon	  	    | 56.92
-----------------------
user_phone	| 555-3434-953 <str/int ?>

*/

func main() {
	// Значения полей могут быть какими угодно - хоть массивами, хоть другми страктами
	
	/*
	type User struct{Name         string
			 		 Address      string
			 		 CardIsActive bool
			 		 Lattitude    float32
			 		 Longitude    float32}
	*/

	//var fomaKinaev User
	//fmt.Println(fomaKinaev)

	
	// var fomaKinaev User = User{"J.Born", "", false, 112.00, 43.98}
	// fmt.Println(fomaKinaev)
	

	/*
	var fomaKinaev User = User{Name: "J.B.",
				     		   CardIsActive: false,
				     		   Lattitude: 112.00,
				     		   Address: "Will Hunting, str, 97",
				     		   Longitude: 43.98}
	fmt.Println(fomaKinaev)
	*/

	/*
	jackBauer := User{Name: "J.B.",
			     	  Address: "FBI Str. 24",
			     	  CardIsActive: false,
			     	  Lattitude: 145.09,
			     	  Longitude: 77.54}
	*/

	//fmt.Println(jackBauer)
	

	//Get element
	
	//fmt.Println(jackBauer.Name)
	//fmt.Println(jackBauer.Lattitude)
	

	//Update
	/*
	fmt.Println(jackBauer.CardIsActive)
	jackBauer.CardIsActive = true
	fmt.Println(jackBauer.CardIsActive)
	*/
	

	// Тип не верный, ошибка
	//jackBauer.CardIsActive = "Yes"
	//jackBauer.CardIsActive = 1
	

	/*
	type Sentence struct{Text string
					     Level int}

	sentence := Sentence{Text: "The Consul watched as Kassad raised the death wand.", 	
						 Level: 1}
	*/
	//fmt.Println(sentence)
	//fmt.Println(sentence.Text)
	
	
	
	type User struct{Username string
					 Level int}

	user := User{Username: "Egor", Level: 1}
	fmt.Println(user)
	fmt.Println(user.Username)
	fmt.Println(user.Level)
	
	
}
