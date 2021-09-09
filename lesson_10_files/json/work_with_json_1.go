package main

import ("encoding/json"
        "fmt"
        "io/ioutil"
        "os"
        )

type Sentence struct {Text string
                      Level int}

func main(){

jsonFile,_ := os.Open("sentences.json")
defer jsonFile.Close()

byteValue, _ := ioutil.ReadAll(jsonFile)
var from_file_sentences []Sentence

json.Unmarshal(byteValue, &from_file_sentences)
fmt.Println(from_file_sentences)


// file, _ := json.MarshalIndent(from_file_sentences, "", " ")
// ioutil.WriteFile("test.json", file, 0644)


}

