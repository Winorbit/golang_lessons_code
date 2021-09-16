package main

import ("encoding/json"
        "fmt"
        "io/ioutil"
        "os")

type Sentence struct {Text string
                      Level int}


func recoveryFunction() {
  if recoveryMessage:=recover(); recoveryMessage != nil {
    fmt.Println(recoveryMessage)
  }
  fmt.Println("This is recovery function...")
}

func getSentences(filename string) []Sentence {
  jsonFile,err := os.Open(filename)
  defer jsonFile.Close()
  //defer recover()

  if err != nil {
    defer recoveryFunction()
    panicMessage := fmt.Sprintf("Function was interrupt because of %s", err.Error())
    panic(panicMessage)
  }

  byteValue, _ := ioutil.ReadAll(jsonFile)
  var extractedSentences []Sentence

  json.Unmarshal(byteValue, &extractedSentences)
  return extractedSentences}


func main() {
  getSentences("unexisted.json")
  //getSentences("sentences.json")
}