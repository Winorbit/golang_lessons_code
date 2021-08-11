package main

import ("fmt"
		"strings"
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
	
	fmt.Println(len(sentences))
	/*
	// answer := 42
	answer := 24
    switch answer {
      case 42: fmt.Println("Oh, hi Douglas!")
      case 24: fmt.Println("Don't recognize reference(")
    };
  */

	/*
	// types mismatch
	answer := "Hello World"
    switch answer {
      case 42: fmt.Println("Oh, hi Douglas!")
      case 24: fmt.Println("Don't recognize reference(")
      default: fmt.Println("Default message")
    };
    */


	/*
	answer := 99
    switch answer {
      case 42: fmt.Println("Oh, hi Douglas!")
      case 24: fmt.Println("Don't recognize reference(")
      default: fmt.Println("Default message")
    };
    */

    /*
	answer := 9
    switch {
      case answer >=42: fmt.Println("Oh, hi Douglas!")
      case answer < 42: fmt.Println("Don't recognize reference(")
      default: fmt.Println("Default message")
    };
    */


    /*
	answer := 9
    switch {
      case answer >=42: fmt.Println("Oh, hi Douglas!")
      break
      case answer < 42: fmt.Println("Don't recognize reference(")
      default: fmt.Println("Default message")
    };
    */

    /*
	mySentence := sentences[3]
	myWord := "sky"
	//myWord := "Sky"
	condition := strings.Contains(mySentence, myWord)

	switch condition{
		case true: fmt.Printf("%v contains in %v\n", myWord, mySentence)
	 	case false: fmt.Println("%v Not contains word %v\n", mySentence, myWord)
	};
	*/

	/*
	mySentence := sentences[3]
	//myWord := "sky"
	myWord := "Sky"
	condition := strings.Contains(strings.ToLower(mySentence), strings.ToLower(myWord))

	switch condition{
		case true: fmt.Printf("%v contains in %v\n", myWord, mySentence)
	};	
	*/
}



