package main

import ("fmt"
		"strings")

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
	
	/*
	fmt.Println(len(sentences))
	
	condition := 2<1
	fmt.Println(condition)

	wordInSentence := strings.Contains("When my time comes \n Forget the wrong that I’ve done.", "done")
	fmt.Println(wordInSentence)
	*/
	
	/*
	if condition{
		fmt.Println("condition is accepted!")
	};

	if wordInSentence{
		fmt.Println("condition is accepted!")
	};

	if condition{
		fmt.Printf("%v \n", condition)
	};
	*/
	

	// if -else
	/*
	if condition{
		fmt.Printf("%v \n", condition)
	}else{
		fmt.Println("Trigered else condition")
	};
	*/

	/*
	value := 10
	if condition{
		fmt.Printf("%v \n", condition)
	} else if value == 10 {
		fmt.Println("Trigered else condition")
	} else {
		fmt.Println("Some unexpected situation")
	};
	*/
	
	
	/*
	mySentence := sentences[4]
	myWord := "across"
	
	condition := strings.Contains(mySentence, myWord)
	if condition{
		fmt.Printf("%v contains in %v\n", myWord, mySentence)
	} else {
		fmt.Println("%v Not contains word %v\n", mySentence, myWord)
	};
	*/
	

	/*
	mySentence := sentences[4]
	myWord := "Across"
	
	condition := strings.Contains(mySentence, myWord)
	if condition{
		fmt.Printf("%v contains in %v\n", myWord, mySentence)
	}else{
		fmt.Printf("%v not contains word %v\n", mySentence, myWord)
	};
	*/
	
	/*
	mySentence := sentences[4]
	myWord := "Across"
	
	condition := strings.Contains(strings.ToLower(mySentence), strings.ToLower(myWord))
	if condition{
		fmt.Printf("%v contains in %v\n", myWord, mySentence)
	}else{
		fmt.Printf("%v not contains word %v\n", mySentence, myWord)
	};
	*/
}



