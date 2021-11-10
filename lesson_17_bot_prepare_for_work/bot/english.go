package bot

import ("fmt"
		    "strings"
		    "os"
		    "io/ioutil"
		    "encoding/json"
		  )


func sentencesFromJson(file_path string) []Sentence {
  jsonFile, err := os.Open(file_path)
  defer jsonFile.Close()

  if err != nil{
    fmt.Println(err.Error())
  }
  byteValue, _ := ioutil.ReadAll(jsonFile)
  var from_file_sentences []Sentence

  json.Unmarshal(byteValue, &from_file_sentences)
  return from_file_sentences
}

func getMatchedSentences(user User, word string, sentences []Sentence)[]string{
	var match_sentences []string
	for _, sentence := range sentences{
	  	wordInSentence := strings.Contains(strings.ToLower(sentence.Text), strings.ToLower(word))
	  	LevelsCompare := user.Level == sentence.Level
		if wordInSentence && LevelsCompare{
			match_sentences = append(match_sentences, sentence.Text)}}
	return match_sentences}


func prepareMessageWithEnglishSentences(foundMatches []string) string {
	message := ""

	// .... //

  return message
}