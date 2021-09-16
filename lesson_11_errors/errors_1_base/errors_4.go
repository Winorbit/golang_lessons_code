package main

import ("encoding/json"
        "fmt"
        "io/ioutil"
        "os"
        "log")

type Sentence struct {Text string
                      Level int}

func getSentences(filename string) []Sentence {
  jsonFile,err := os.Open(filename)
  defer jsonFile.Close()

  if err != nil {
    errorMessage := fmt.Sprintf("Function was failed because of %s", err.Error())
    log.Print(errorMessage)
  }

  byteValue, _ := ioutil.ReadAll(jsonFile)
  var extractedSentences []Sentence

  json.Unmarshal(byteValue, &extractedSentences)
  return extractedSentences}


func main() {
  //log.Print("just message")
  getSentences("unexisted.json")
  //getSentences("sentences.json")
}