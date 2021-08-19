package main

import "fmt"

func main() {
	
	type Sentence struct{Text string
					     Level int}
	/*
	type User struct{Username string
					 Level int}
	*/

	// Обявление нового объекта из struct - Вариант 1
	/*
	sentence := Sentence{Text: "The Consul watched as Kassad raised the death wand.", 	
						 Level: 1}
	fmt.Println(sentence)
	*/


	/*
	Обявление нового объекта из struct - Вариант 2
	Это выделит память для всех полей, 
	присвоит каждому из них нулевое значение и вернет указатель (*Sentence)
	*/

	/*
	newSentence := new(Sentence)
	fmt.Println(newSentence)
	newSentence.Text = "Hello Dolly"
	fmt.Println(newSentence)
	newSentence.Level = 1
	fmt.Println(newSentence)
	*/
	
	/*
	Обявление нового объекта из struct - Вариант 3
	Подобно другим типами данных, будет создана локальная переменная типа Sentence,
	чьи поля по умолчанию будут равны нулю
	(0 для int, 0.0 для float, "" для string, nil для указателей, етс.). 
	*/

	/*
	var newSentence Sentence
	fmt.Println(newSentence)
	newSentence.Text = "Hello Dolly"
	fmt.Println(newSentence)
	newSentence.Level = 1
	fmt.Println(newSentence)
	*/
	

	/*
	type User struct{Username string
					 Level int
					 FavoriteThemes []string}

	andrew := User{Username: "Andrew",
			   	   Level: 2,
			   	   FavoriteThemes: []string{"sci-fi", "music", "videogames"}}

	//fmt.Println(andrew)
	//fmt.Println(andrew.FavoriteThemes)
	
	fmt.Println(andrew.FavoriteThemes[1])
	andrew.FavoriteThemes[1] = "jenga"
	fmt.Println(andrew.FavoriteThemes[1])
	*/
	

	type User struct{Username string 		 `json:"username"`
					 Level int 			     `json:"level"`
					 FavoriteThemes []string `json:"favorite_themes"`}
	/*
	andrew := User{Username: "Andrew",
			   	   Level: 2,
			   	   FavoriteThemes: []string{"sci-fi", "music", "videogames"}}
	*/

	
	//fmt.Println(andrew)
	//fmt.Println(andrew.FavoriteThemes)
	//fmt.Println(andrew.Username)
	//fmt.Println(&andrew)

	type User2 struct{Id int}

	var oleg *User2 = new(User)
	fmt.Println(&oleg)
	//fmt.Println(oleg)

	//fmt.Println(oleg)
	//oleg = &andrew
	//fmt.Println(oleg)	


}
