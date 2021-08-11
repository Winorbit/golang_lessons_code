package main

import ("fmt"
		  //"strings"
		)

func main() {
	sentences := []string{"When my time comes \n Forget the wrong that I’ve done.",
				 				  			  "In a hole in the ground there lived a hobbit.",
				 				  			  "Airplanes are beautiful, cursed dreams, waiting for the sky to swallow them up.", 
				 				  			  "The sky the port was the color of television, tuned to a dead channel.",
				 				  			  "I love the smell of napalm in the morning.",
				 				  			  "The man in black fled across the desert, and the gunslinger followed.",
				 				  			  "The Consul watched as Kassad raised the death wand.",
				 				  			  "If you want to make enemies, try to change something.",
				 				  			  "Of course, they say every atom in our bodies was once part of a star. Maybe I'm not leaving... maybe I'm going home",
				 				  			  "We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore",
				 				  			  "I learned very early the difference between knowing the name of something and knowing something.",
				 				  			  "On the eastern horizon, a huge cloud of smoke from burning oil tanks stretched across the sky.",
			 								}


	/* Описание цикла, подготовка

	итератор  условие, когда прекращать, что делать на каждом шаге{
		// Тело. Здесь статует цикл
	}
	*/

	for x := 1; x < 10; x++ {
	  fmt.Println(x)
	}


	/* 
	Можно добавить только условие и значение переменной, 
	но если переменная не изментся, то мы попадаем в вечный цикл.
	*/
	for x := 3; x < 10{
    	fmt.Println(x)
	}

	// Переменную можно задать и в другом месте.
	x := 3
	for x < 10{
    	fmt.Println(x)
	}

	/*
	x := 3
	for x < 10{
    	fmt.Println(x)
    	x += 2}

    */

    // С помощью этого же цикла мы, в принципе, можем и перебирать массивы.

	myStrings := []string{"first", "second", "third", "fourth"}

	for x := 1; x < 4; x++ {
	  fmt.Println(myStrings[x])
	}


	/* Можно словить OOR
	for x := 1; x =< 4; x++ {
	  fmt.Println(sentences[x])
	}

	for x := 1; x < 10; x++ {
	  fmt.Println(sentences[x])
	}
	*/

	/*
	Чтоб этого не было, и не забивать вручную длинну, можно сделать так.
	for x := 1; x < len(sentences); x++ {
	  fmt.Println(sentences[x])
	}
	*/


	//Range отдает два значения - index, value.
	/*
	for index, value := range sentences{ 
		fmt.Println(index, value)}
	*/


	// Declared and not used
	/*
	for index, value := range sentences{ 
		fmt.Println(value)}
	*/

	/* 
	// Решим проблему с помощью _
	Такой прикол и в питоне доступен.
	for _, value := range sentences{ 
		fmt.Println(value)}

	*/

	// Можно и в одну строку.
	//for _, value := range sentences{fmt.Println(value)}
}



