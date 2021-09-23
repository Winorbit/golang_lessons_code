package main

import "net/http"
import "encoding/json"
import "fmt"

/*
Чтобы извлечь данные о последних сообщениях боту, нам нужно извлчеь информацию в JSON.
А для этого JSON подготовить соответствующие структуры, куда упакуем данные, используя нужные нам поля
*/

type Chat struct {
  Id int `json:"id"`
}

type Message struct {
  Text     string   `json:"text"`
  Chat     Chat     `json:"chat"`
}

type Update struct {
  UpdateId int     `json:"update_id"`
  Message  Message `json:"message"`
}

type Updates struct {
  Updates []Update  `json:"result"`
}

const rootURL = "https://api.telegram.org/bot"
var token = "2032676340:AAGkrcOu-qdN8f0apzXJQuqm0KxoEs0Tjik"

func main(){
  updatesURL := fmt.Sprintf("%s%s/getUpdates", rootURL, token)
  res, err := http.Get(updatesURL)
  defer res.Body.Close()

  if err != nil {
    fmt.Println(err.Error())
  }

  var updates Updates
  decoder := json.NewDecoder(res.Body)
  decoder.Decode(&updates)

  lastMessageIndex := len(updates.Updates) - 1 
  lastMessageText := updates.Updates[lastMessageIndex].Message.Text
  chatId := updates.Updates[lastMessageIndex].Message.Chat.Id

  fmt.Println(lastMessageText, chatId)
}
