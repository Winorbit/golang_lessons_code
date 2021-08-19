package main

import (
	"fmt"
	"strings"
)

type Sentence struct {
	Text  string
	Level int
}

var quotes []Sentence = []Sentence{
	{Text: "When my time comes \n Forget the wrong that I’ve done.", Level: 1},
	{Text: "In a hole in the ground there lived a hobbit.", Level: 1},
	{Text: "Airplanes are beautiful, cursed dreams, waiting for the sky to swallow them up.", Level: 2},
	{Text: "The sky the port was the color of television, tuned to a dead channel.", Level: 2},
	{Text: "I love the smell of napalm in the morning.", Level: 1},
	{Text: "The man in black fled across the desert, and the gunslinger followed.", Level: 2},
	{Text: "The Consul watched as Kassad raised the death wand.", Level: 1},
	{Text: "If you want to make enemies, try to change something.", Level: 1},
	{Text: "Of course, they say every atom in our bodies was once part of a star. Maybe I'm not leaving... maybe I'm going home", Level: 3},
	{Text: "We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore", Level: 1},
	{Text: "I learned very early the difference between knowing the name of something and knowing something.", Level: 2},
	{Text: "On the eastern horizon, a huge cloud of smoke from burning oil tanks stretched across the sky.", Level: 3},
}

func main() {

	// myWord := "smoke"
	myWord := "across"
	// myWord := "libriary"

	var match_sentences []string
	fmt.Println(match_sentences)
	for _, sentence := range quotes {
		wordInSentence := strings.Contains(strings.ToLower(sentence.Text), strings.ToLower(myWord))
		if wordInSentence {
			match_sentences = append(match_sentences, sentence.Text)
		}
	}
	// fmt.Println(match_sentences)

	sentenceCount := len(match_sentences)

	// if len(match_sentences) == 0 {
	// 	fmt.Println("Sorry, no matches found for your word =(")
	// }
	// if len(match_sentences) == 1 {
	// 	fmt.Println("There's one sentence found for your word: \n\n")
	// 	fmt.Println(match_sentences[0])
	// }
	// if len(match_sentences) > 1 {
	// 	for _, predlozheniye := range match_sentences {
	// 		fmt.Println(predlozheniye)
	// 	}
	// }

	switch {
	case sentenceCount == 0:
		fmt.Println("Sorry, no matches found for your word =(")
	case sentenceCount == 1:
		fmt.Println("There's one sentence found for your word: \n\n")
		fmt.Println(match_sentences[0])
	case sentenceCount > 1:
		for _, predlozheniye := range match_sentences {
			fmt.Println(predlozheniye)
		}
	}
}
