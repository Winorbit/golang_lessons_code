package bot

import ("net/http"
        "encoding/json"
        "fmt"
        "net/url"
        "strconv"
      )

func Contains(elements []string, element string) bool {
  for _, el := range elements {
    if el == element {
      return true
    }
  }
  return false
}


func getUpdates(token string) Updates {
    updatesURL := fmt.Sprintf("%s%s/getUpdates", rootURL, token)
    var updates Updates

    res, err := http.Get(updatesURL)
    statusCode := res.Status
    defer res.Body.Close()

    if err != nil {
      fmt.Println(err.Error())
    }
    
    if Contains(okCodes, statusCode) == false {
      fmt.Printf("Request failed with status-code: %s", statusCode)
    }
    
    decoder := json.NewDecoder(res.Body)
    decoder.Decode(&updates)
    return updates
}



func sendMessage(chatId string, message string, token string) {
  sendMessageURL := fmt.Sprintf("%s%s/sendMessage", rootURL, token)
  res, err := http.PostForm(sendMessageURL,
                            url.Values{"chat_id": {chatId},
                                       "text": {message} })
  statusCode := res.Status[0:3]

  if err != nil {
    fmt.Println(err.Error())
  }

  if Contains(okCodes, statusCode) == false {
    fmt.Printf("Request failed with status-code: %s", statusCode)}
  }


func processMessages(chatId string, message string, token string){
  if message == "Hello"{
    sendMessage(chatId, "Your welcome!", token)
  } else {
    sendMessage(chatId, message, token)
  }
}


func Pooling(token string){
  lastMessageId := 0 

  for {
    updates := getUpdates(token)
    lastUpdateIndex := len(updates.Updates) - 1 
    MessageId := updates.Updates[lastUpdateIndex].Message.Id
    
    if MessageId > lastMessageId{
      chatId := strconv.Itoa(updates.Updates[lastUpdateIndex].Message.Chat.Id)
      lastMessageText := updates.Updates[lastUpdateIndex].Message.Text
      processMessages(chatId, lastMessageText, token)
      lastMessageId = MessageId
    }
  }
}
