package values

import ("encoding/json"
        "io/ioutil"
        "fmt"
        "log"
        "os")

var EnglishLevels map[int]string = map[int]string{1: "Beginer",
												                          2: "Middle",
												                          3: "Intermediate"}

func GetSentences(filename string) []Sentence {
  jsonFile,openErr := os.Open(filename)
  defer jsonFile.Close()

  if openErr != nil {
    log.Print("Problem with opening file: %s", openErr.Error())
  }

  byteValue, byteErr := ioutil.ReadAll(jsonFile)
  if byteErr != nil {
    errorMessage := fmt.Sprintf("Problem with reading: %s", byteErr.Error())
    log.Print(errorMessage)

  }

  var fromFileSentences []Sentence
  json.Unmarshal(byteValue, &fromFileSentences)
  return fromFileSentences
}