package main

import "fmt"

func main() {
	// Если вместо длинны указать [...], то длинна = кол-во элементов
	/*
	var sentences = [...]string{"When my time comes \n Forget the wrong that I’ve done.",
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


	// Вот так не получится
	// sentences[15] = "Brand new element"


	// Можно задавать элементу номер его индекса вручную
	/*
	weiredList := [...]string{12: "12 Monkeys", 3: "Third man"}
	fmt.Println(weiredList)
	fmt.Println(weiredList[12])
	fmt.Println(weiredList[3])
	fmt.Println(weiredList[0])
	*/
	
	sentences := []string{"When my time comes \n Forget the wrong that I’ve done.",
				 		  "In a hole in the ground there lived a hobbit.",
				 		  "The sky the port was the color of television, tuned to a dead channel.",
				 		  "I love the smell of napalm in the morning.",
				 		  "The man in black fled across the desert, and the gunslinger followed.",
				 		  "The Consul watched as Kassad raised the death wand.",
				 		  "If you want to make enemies, try to change something.",
				 		  "We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore",
				 		  "I learned very early the difference between knowing the name of something and knowing something.",
			 			}
	

	// DELETE

	
	fmt.Println(len(sentences))
	sentences[3] = sentences[len(sentences)-1] 
	sentences[len(sentences)-1] = "" 
	sentences = sentences[:len(sentences)-1]
	fmt.Println(len(sentences))
	

	// Пример на более маленьком срезе
	
	var arr1 []int = []int{1,2,3,4,5,6,7}
	fmt.Println(arr1)
	copy(arr1[3:], arr1[4:])
	fmt.Println(arr1)
	arr1 = arr1[:len(arr1)-1]
	fmt.Println(arr1)

	
}
