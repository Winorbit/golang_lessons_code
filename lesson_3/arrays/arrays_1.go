package main

import ("fmt"
	   "reflect")


/*

var (sentence1 string = "The sky the port was the color of television, tuned to a dead channel."
     sentence2 = "The man in black fled across the desert, and the gunslinger followed."
     sentence3 ="We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore"
     sentence4 = "In a hole in the ground there lived a hobbit."
 )


sentence1 := "The sky the port was the color of television, tuned to a dead channel."
sentence2 := "The man in black fled across the desert, and the gunslinger followed."
sentence3 :="We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore"
sentence4 := "In a hole in the ground there lived a hobbit."
*/

func main() {
	// Длинна и типы нужны для выделения памяти под массив
	var words [3]string = [3]string{"hello", "world","!"}
	fmt.Println(words)

	// Вот так можно
	// var nums [3]int = [3]int{8, 15, 16}

	//И так можно. Недостающий элемент будет пустым
	// var nums [3]int = [3]int{15, 16}
	// fmt.Println(nums)
 
	// Так - уже нет
	// var nums [3]int = [3]int{"8", 15, 16}

	// var nums [3]int = [3]int{15, 16, 23, 42}


	/*
	var sentencesArray [10]string = [10]string{"When my time comes \n Forget the wrong that I’ve done.",
				 					   "In a hole in the ground there lived a hobbit.",
				 					   "The sky the port was the color of television, tuned to a dead channel.",
				 					   "I love the smell of napalm in the morning.",
				 					   "The man in black fled across the desert, and the gunslinger followed.",
				 					   "The Consul watched as Kassad raised the death wand.",
				 					   "If you want to make enemies, try to change something.",
				 					   "We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore",
				 					   "I learned very early the difference between knowing the name of something and knowing something.",
			 						  }
	*/


	// Плохой нейминг. Array тут ни к чему. просто sentences

	/*
	var sentences [10]string = [10]string{"When my time comes \n Forget the wrong that I’ve done.",
				 "In a hole in the ground there lived a hobbit.",
				 "The sky the port was the color of television, tuned to a dead channel.",
				 "I love the smell of napalm in the morning.",
				 "The man in black fled across the desert, and the gunslinger followed.",
				 "The Consul watched as Kassad raised the death wand.",
				 "If you want to make enemies, try to change something.",
				 "We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore",
				 "I learned very early the difference between knowing the name of something and knowing something.",
			 }

	*/

	// Можно и так объявить
	sentences := [10]string{"When my time comes \n Forget the wrong that I’ve done.",
				 		"In a hole in the ground there lived a hobbit.",
				 		"The sky the port was the color of television, tuned to a dead channel.",
				 		"I love the smell of napalm in the morning.",
				 		"The man in black fled across the desert, and the gunslinger followed.",
				 		"The Consul watched as Kassad raised the death wand.",
				 		"If you want to make enemies, try to change something.",
				 		"We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore",
				 		"I learned very early the difference between knowing the name of something and knowing something.",
			 }

	// Узнать тип
	fmt.Println(reflect.TypeOf(sentences))
	// Узнать длинну
	fmt.Println(len(sentences))

	//GET element
	fmt.Println(sentences[0])
	fmt.Println(sentences[3])
	// Отрицатеьный индекс не работает
	//fmt.Println(sentences[-1])
	fmt.Println(sentences[len(sentences)-1])
	
	//index out of range
	//sentences[99]

	//GET slice
	fmt.Println(sentences[1:3])
	fmt.Println(reflect.TypeOf(sentences[1:3]))
	firstThreeElements := sentences[:3]
	fmt.Println(firstThreeElements)
	lastThreeElements := sentences[3:]
	fmt.Println(lastThreeElements)

	//UPDATE
	fmt.Println(sentences[5])
	sentences[5] = "My new sentence"
	fmt.Println(sentences[5])
	//index out of range
	//sentences[99] = "This is new value"


	// Slices
	testSlice := sentences[:5]
	//fmt.Println(testSlice[-1], len(testSlice))
	fmt.Println(testSlice[len(testSlice)-1], len(testSlice))

	testSlice = append(testSlice, "Appended value")
	fmt.Println(testSlice[len(testSlice)-1], len(testSlice))

	var newSlice []string
	fmt.Println(newSlice, reflect.TypeOf(newSlice))
	newSlice = append(testSlice, "Appended value")
	fmt.Println(newSlice)

	// Окей но у нас есть ограничение длинны. Для чего это удобно?

	/*
	okCodes := [6]int{200,201,202,203,204,205}
	fmt.Println(okCodes)
	*/

	okCodes := [10]int{200,201,202,203,204,205}
	fmt.Println(okCodes)

	/* Можно и так

	bools := [10]bool{true}
	fmt.Println(bools)
	bools[7] = true
	fmt.Println(bools)
	bools[12] = true
	*/
}
