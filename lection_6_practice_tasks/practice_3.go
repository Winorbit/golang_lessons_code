package main

import (
	"fmt"
	str "strings"
)

type Sentence struct {
	Text  string
	Level int
}

type User struct {
	Username string
	Level    int
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

	me := User{Username: "Egor", Level: 3}
	//myWord := "smoke"
	myWord := "across"
	//myWord := "libriary"
	var match_sentences []string

	for _, sentence := range quotes {

		wordInSentence := str.Contains(str.ToLower(sentence.Text), str.ToLower(myWord))
		LevelsCompare := me.Level >= sentence.Level
		if wordInSentence && LevelsCompare {
			match_sentences = append(match_sentences, sentence.Text)
		}
	}

	if len(match_sentences) == 0 {
		fmt.Println("Не найдено подходящих предложений.")
	}
	if len(match_sentences) == 1 {
		fmt.Println(match_sentences[0])
	}
	if len(match_sentences) > 1 {
		for _, sentence := range match_sentences {
			fmt.Println(sentence)
		}
	}

}
