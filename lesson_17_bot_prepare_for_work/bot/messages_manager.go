package bot

import ("net/http"
        "encoding/json"
        "fmt"
        "net/url"
        "log"
        "os"
        "strconv"
      )

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

var testUser = User{Level: 2}


func Contains(elements []string, element string) bool {
  for _, el := range elements {
    if el == element {
      return true
    }
  }
  return false
}

func messageIsInt(message string) bool {
  //......///
}


func getUpdates(token string) Updates {
    updatesURL := fmt.Sprintf("%s%s/getUpdates", rootURL, token)
    var updates Updates

    res, err := http.Get(updatesURL)
    // infoLog.Printf("request for BotUpdates send to %v", updatesURL)
    statusCode := res.Status[0:3]
    defer res.Body.Close()

    if err != nil {
      errorLog.Println(err.Error())
    }
    if !Contains(okCodes, statusCode){
      errorLog.Printf("Request failed with status-code: %s", statusCode)
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
    errorLog.Println(err.Error())
  }
  if Contains(okCodes, statusCode) {
    infoLog.Printf("Message %v successfully send with status-code: %s", message, statusCode)
  }
  if !Contains(okCodes, statusCode) {
    errorLog.Printf("Sending message %v to userfailed with status-code: %s", message, statusCode)
  }
}


func processMessages(chatId string, message string, token string){
  if Contains(botCommands, message){
    if message == "/start"{
      // .....
    }
  }
  if !messageIsInt(message){
      matches :=  //......
      messageToUser := prepareMessageWithEnglishSentences(matches)
      sendMessage(//....///)
    }
}


func Pooling(token string){
  lastMessageId := 0 
  for {
    updates := getUpdates(token)
    // ....
    if len(updates.Updates) != 0{
      lastUpdateIndex := len(updates.Updates) - 1 
      messageId := updates.Updates[lastUpdateIndex].Message.Id

      if messageId > lastMessageId{
        chatId := strconv.Itoa(updates.Updates[lastUpdateIndex].Message.Chat.Id)
        lastMessageText := updates.Updates[lastUpdateIndex].Message.Text
        processMessages(chatId, lastMessageText, token)
        lastMessageId = messageId
      }
    }
  }
}
