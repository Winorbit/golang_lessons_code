package main

import ("fmt"
		"strings")

type Sentence struct {Text string
		    		  Level int}

quotes:= []Sentence{
					{Text: "When my time comes \n Forget the wrong that I’ve done.", Level: 1},
				 	{Text: "In a hole in the ground there lived a hobbit.", Level: 1},
				 	{Text: "Airplanes are beautiful, cursed dreams, waiting for the sky to swallow them up.",  Level: 2},
				 	{Text: "The sky the port was the color of television, tuned to a dead channel.", Level:2 },
				 	{Text: "I love the smell of napalm in the morning.", Level:1 },
				 	{Text: "The man in black fled across the desert, and the gunslinger followed.", Level:2 },
				 	{Text: "The Consul watched as Kassad raised the death wand.", Level: 1},
				 	{Text: "If you want to make enemies, try to change something.", Level:1 },
				 	{Text: "Of course, they say every atom in our bodies was once part of a star. Maybe I'm not leaving... maybe I'm going home", Level:3},
				 	{Text: "We're not gonna take it. \n Oh no, we ain't gonna take it \nWe're not gonna take it anymore", Level: 1 },
				 	{Text: "I learned very early the difference between knowing the name of something and knowing something.", Level: 2 },
				 	{Text: "On the eastern horizon, a huge cloud of smoke from burning oil tanks stretched across the sky.", Level: 3 },
				 	}

func main() {

	myWord := "smoke"
	for _, sentence := range quotes{
		...
	};


}