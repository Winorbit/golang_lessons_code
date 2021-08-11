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
	fmt.Println(len(sentences))
	/*
	myWord := "across"
	for x := 1; x < len(sentences); x++ {
	  sentence := sentences[x]
	  wordInSentence := strings.Contains(sentence, myWord)
	  if wordInSentence {
	    fmt.Println(sentence)
	  }
	};
	*/

	/*
	myWord := "across"
	for x := 1; x < len(sentences); x++ {
	  sentence := sentences[x]
	  wordInSentence := strings.Contains(sentence, myWord)
	  if wordInSentence {
	    fmt.Println(sentence)
	    break
	  }
	};
	*/

    for i := 0; i < 10; i++ {
        if i == 5 {
            fmt.Println("Continuing loop")
            continue // break here
            fmt.Println("This not will be printed")
        }
        fmt.Println("The value of i is", i)
    };

	/*
	myWord := "across"
	for x := 1; x < len(sentences); x++ {
	  sentence := sentences[x]
	  condition := strings.Contains(strings.ToLower(sentence), strings.ToLower(myWord))
	  if condition {fmt.Println(sentence)}
	};
	*/


   /*
   var matchSentences []string
	//myWord := "sky"
	myWord := "across"

	fmt.Println(matchSentences)

	for _, sentence := range sentences{
	  	condition := strings.Contains(strings.ToLower(sentence), strings.ToLower(myWord))
		if condition{
			matchSentences = append(matchSentences, sentence)}}

   
   fmt.Println(matchSentences)
   */
}



