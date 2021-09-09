package main

import ("encoding/json"
        f "fmt"
        //"io/ioutil"
        //"os"
        )

type Sentence struct{Text string
                     Level int
                    }

func main(){
// Базовый пример
/*
simple_sentence := Sentence{Text: "Lorem ipsum", Level: 0}
json_simple_sentence, err := json.Marshal(simple_sentence)

if err != nil {
    f.Println(err.Error())
    }
    
f.Println(json_simple_sentence)
f.Println(string(json_simple_sentence))
*/

// Marshaling man sentences
/*
simple_sentence := Sentence{Text: "Lorem ipsum", Level: 0}
sentences := []Sentence{
                        {Level: 0, Text: "In a hole in the ground there lived a hobbit."},
                        simple_sentence,
                        {Text: "The sky the port was the color of television, tuned to a dead channel.", Level:1},
                    }

json_sentences, _ := json.Marshal(sentences)
f.Println(json_sentences)
f.Println(string(json_sentences))
*/


// Из JSON в STRUCT
// ` а не '
/*
json_sentence := []byte(`{"text": "If you want to make enemies, try to change something.", 
                                  "level": 2}`)
// создаем структуру, в которую это декодируем

var go_sentence Sentence
f.Println(go_sentence)
*/
/*
Возвращает только ошибку, т.к мы используем указатель 
и не создаем новое значение ,а меняем значение прежней переменной
*/

//err := json.Unmarshal(json_sentence, &go_sentence)

// json.Unmarshal(json_sentence, &go_sentence)
// f.Println(go_sentence)


// Можно как убрать поле из JSON или структуры
var go_sentences []Sentence
jsoned_sentences := []byte(`[
                             {"level": 0, "text": "In a hole in the ground there lived a hobbit."},
                             {"text": "The sky the port was the color of television, tuned to a dead channel.", "level":1}
                            ]`)
// Запятая

json.Unmarshal(jsoned_sentences, &go_sentences)
//f.Println(go_sentences)

for x := range go_sentences {
    f.Println(go_sentences[x])
}

}

