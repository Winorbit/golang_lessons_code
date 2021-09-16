package main

import ("encoding/json"
        "fmt"
        "io/ioutil"
        "os"
        )
/*
В Go нет специального механизма обработки исключений, идея здесь, как и всегда, 
в смотоятельной сборке механизмов под себя.
*/

type Sentence struct {Text string
                      Level int}

/*
func getSentences(filename string) ([]Sentence, error){
  jsonFile,err := os.Open(filename)
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)
  var extractedSentences []Sentence

  json.Unmarshal(byteValue, &extractedSentences)
  return extractedSentences, err}
*/

func getSentences(filename string) []Sentence {
  jsonFile,err := os.Open(filename)
  defer jsonFile.Close()

  //В Go ошибка это просто тип данных - https://golang.org/ref/spec#Errors
  if err != nil {
    fmt.Println(err.Error())
  }

  byteValue, _ := ioutil.ReadAll(jsonFile)
  var extractedSentences []Sentence

  json.Unmarshal(byteValue, &extractedSentences)
  return extractedSentences}


func main() {
  //res := getSentences("sentences.json")
  fmt.Println(getSentences("sentences.json"))}