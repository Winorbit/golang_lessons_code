package main

import "net/http"
import "fmt"
import "io"
import "net/url"



func main(){

  /*
  botInfoURL := "https://api.telegram.org/bot1818081328:AAFAPDBGdpbhEN5qEgeI6uzaTZFxmaRxyGs/getme"
  resp, _ := http.Get(botInfoURL)

  body := resp.Body
  fmt.Println(body)

  bodyData, _ := io.ReadAll(body)
  fmt.Println(bodyData)
  fmt.Println(string(bodyData))
  */


  // send message via GET
  /*
  sendMessageURL := "https://api.telegram.org/bot1818081328:AAFAPDBGdpbhEN5qEgeI6uzaTZFxmaRxyGs/sendMessage?chat_id=565281269=&text=hellllllo "
  res, err := http.Get(sendMessageURL)
  if err =! nil {
	  fmt.Println(err.Error()
  }
  */

  // send message via POST
  /*
  sendMessageURL := "https://api.telegram.org/bot1818081328:AAFAPDBGdpbhEN5qEgeI6uzaTZFxmaRxyGs/sendMessage"

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
