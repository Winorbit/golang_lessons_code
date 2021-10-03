package main

import "net/http"
import "encoding/json"
import "fmt"
import "net/url"

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

func getUpdates(token string) Updates {
  updatesURL := fmt.Sprintf("%s%s/getUpdates", rootURL, token)
  res, err := http.Get(updatesURL)
  defer res.Body.Close()

  if err != nil {
    fmt.Println(err.Error())
  }
  var updates Updates
  decoder := json.NewDecoder(res.Body)
  decoder.Decode(&updates)
  return updates}


func sendMessage(chatId string, message string, token string) {
  sendMessageURL := fmt.Sprintf("%s%s/sendMessage", rootURL, botToken)
  _, err := http.PostForm(sendMessageURL,
                          url.Values{"chat_id": {chatId},
                                     "text": {message} })
  if err != nil {
    fmt.Println(err.Error())
  }
  // Наша функция не отработает, и, при этом, никак не уведомит нас, если статус ответа будет на 2ХХ
}

func main(){
  updates := getUpdates(botToken)
  lastMessageIndex := len(updates.Updates) - 1 
  chatId := updates.Updates[lastMessageIndex].Message.Chat.Id

  // Отправить свое сообщение
  sendMessage(string(chatId), "Hello!", botToken)

  // Отправить эхо
  echoMessage := updates.Updates[lastMessageIndex].Message.Text
  sendMessage(string(chatId), echoMessage, botToken)
}
