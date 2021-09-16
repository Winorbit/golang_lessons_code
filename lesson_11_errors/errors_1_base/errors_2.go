package main

import ("encoding/json"
        "fmt"
        "io/ioutil"
        "os")

type Sentence struct {Text string
                      Level int}

/*
Для надежной работы Go предлагает такое комбо:
error - тип данных
defer - объявляет, что функция будет выполнена в любом случае.
panic - вызывает искуственное пркрывание программы
recover - позволяяет перехватить прерывание и вернуть контроль
*/

func getSentences(filename string) []Sentence {
  jsonFile,err := os.Open(filename)
  defer jsonFile.Close()

  if err != nil {
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