package main

import ("fmt"
		"strings"
		"github.com/k0kubun/pp"
		"packages_with_main/src/values")

func getMatchedSentences(user values.User, word string, sentences []values.Sentence)[]string{
	var match_sentences []string
	for _, sentence := range sentences{
	  	wordInSentence := strings.Contains(strings.ToLower(sentence.Text), strings.ToLower(word))
	  	LevelsCompare := user.Level == sentence.Level
		if wordInSentence && LevelsCompare{
			match_sentences = append(match_sentences, sentence.Text)}}
	return match_sentences}

func createUserProfile(username string, level int) values.User{
	new_user := values.User{Username: username,
				 	 Level: level}
	return new_user}

func sendSentencesToUser(user values.User, word string, sentences []values.Sentence){
    matches := getMatchedSentences(user, word, sentences) 
    matches_count := (len(matches))
    switch {case matches_count  == 0:
	        	fmt.Println("Sorry, nothing on your request.")
	    	case matches_count == 1:
				fmt.Println(matches[0])
            case matches_count > 1:
				for x := 0; x < matches_count; x++ {
		    		fmt.Println(matches[x]) }}}

func Dialog(){
	var username string
    var englishLevel int
    var word string

    fmt.Println("Привет! Как к тебе обращаться?")
    fmt.Scanf("%s", &username)

    fmt.Println("Какой у тебя уровень английского? Выбери номер варианта:\n")
    pp.Print(values.EnglishLevels)
    fmt.Scanf("%d", &englishLevel)

    fmt.Println("Какое бы ты слово хотел выучить?")
    fmt.Scanf("%s", &word)

    user := createUserProfile(username, englishLevel)
    sendSentencesToUser(user, word, values.Quotes)}
