package main

import ("net/http"
        "fmt"
        "net/url")

// Сделаем URL более аккуратным. Корневой сделаем константой
const rootURL = "https://api.telegram.org/bot"
// А токен, пока, просто переменной
var token = "2032676340:AAGkrcOu-qdN8f0apzXJQuqm0KxoEs0Tjik"

func main(){
  
  botInfoURL := fmt.Sprintf("%s%s/getMe", rootURL, token)
  resp, _ := http.Get(botInfoURL)
  body := resp.Body
  fmt.Println(body)

  bodyData, _ := io.ReadAll(body)
  fmt.Println(bodyData)
  fmt.Println(string(bodyData))
  
  // Отправка сообщения методом GET
  /*
  sendMessageURL := fmt.Sprintf("%s%s/sendMessage?chat_id=565281269=&text=hello", rootURL, token)
  _, err := http.Get(sendMessageURL)
  if err != nil {
	  fmt.Println(err.Error())
  }
  */

  // Отправка сообщения методом POST
  /*
  sendMessageURL := fmt.Sprintf("%s%s/sendMessage", rootURL, token)
  _, err := http.PostForm(
		sendMessageURL,
		url.Values{
			"chat_id": {"565281269"},
			"text":    {"hello"},
		})
  if err != nil {
	  fmt.Println(err.Error())
  }
  */

}
