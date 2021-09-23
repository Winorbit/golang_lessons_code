package main

import ("net/http"
        "fmt"
        "net/url")

const rootURL = "https://api.telegram.org/bot"
var token = "2032676340:AAGkrcOu-qdN8f0apzXJQuqm0KxoEs0Tjik"

func sendMessage(chatId string, message string){
  sendMessageURL := fmt.Sprintf("%s%s/sendMessage", rootURL, token)
  http.PostForm(sendMessageURL,
			          url.Values{"chat_id": {chatId},
					                 "text": {message} })}

func main(){
  sendMessage("565281269", "Hello!")
}
