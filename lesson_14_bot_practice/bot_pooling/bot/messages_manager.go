package bot

import ("net/http"
        "encoding/json"
        "fmt"
        "net/url"
        "strconv"
      )

func contains(elements []string, element string) bool {
  for _, el := range elements {
    if el == element {
      return true
    }
  }
  return false
}


func getUpdates(token string) Updates {
  ...
}



func sendMessage(chatId string, message string, token string) {
  sendMessageURL := fmt.Sprintf("%s%s/sendMessage", rootURL, token)
  res, err := http.PostForm(sendMessageURL,
                            url.Values{"chat_id": {chatId},
                                       "text": {message} })
  statusCode := res.Status[0:3]

  ...
  }


func processMessages(chatId string, message string, token string){
  if message == "Hello"{
    sendMessage(chatId, "Your welcome!", token)
  } else {
    sendMessage(chatId, message, token)
  }
}


func Pooling(token string){
  ...
}
