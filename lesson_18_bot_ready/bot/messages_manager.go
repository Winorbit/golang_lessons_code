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


func numberIsEnglishLevel(newLevel int, availableLevels []int) bool {
  // ...
}


func messageIsInt(message string) bool {
  if _, err := strconv.Atoi(message); err == nil {
    return true
  } else {
    return false
  }
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


func addUserToSession(user User){
  // ...
}


func processMessages(user_id int, chatId string, message string, token string){
  userFromSession := sessions[user_id]

  if Contains(botCommands, message){
    if message == "/start"{
      sendMessage(chatId, startMessage, token)
    }
    if message == "/start_learning" || message == "/change_level"{
      sendMessage(chatId, "Let's do this!!", token)
      sendMessage(chatId, updateEnglishLevelMessage, token)
      userFromSession.UpdateInProgress = true
    }
    if message == "/my_level"{
      // ...
    }
  }

  if !Contains(botCommands, message){
    if messageIsInt(message){
      intMessage, err := strconv.Atoi(message)
      if err != nil{
        fmt.Println(err.Error())
      }
      if numberIsEnglishLevel(intMessage, availableLevels){
        if userFromSession.UpdateInProgress{
          userFromSession.Level = intMessage
          sendMessage(chatId, "Your level  successfully updated!", token)
          infoLog.Printf("Level for user %v was successfully updated from to %v", userFromSession, intMessage)
        } else{
          sendMessage(chatId, "Sorry, but i can't understand", token)
          infoLog.Printf("Failed updating user level for user %v with message %v", userFromSession, message)
        } 
      } else {
        sendMessage(chatId, "Sorry, but i can't understand", token)
        infoLog.Printf("Can't send correct answer to user %v on message: %v", userFromSession, message)
      }
    }
    if !messageIsInt(message){
      matches := getMatchedSentences(userFromSession, message, englishSentences)
      messageToUser := prepareMessageWithEnglishSentences(matches)
      sendMessage(chatId, messageToUser, token)
    }
  }
}


func Pooling(token string){
  lastMessageId := 0 
  for {
    updates := getUpdates(token)
    if len(updates.Updates) != 0{
      lastUpdateIndex := len(updates.Updates) - 1 
      messageId := updates.Updates[lastUpdateIndex].Message.Id
      currentUser := updates.Updates[lastUpdateIndex].Message.User
      addUserToSession(currentUser)

      if messageId > lastMessageId{
        chatId := strconv.Itoa(updates.Updates[lastUpdateIndex].Message.Chat.Id)
        lastMessageText := updates.Updates[lastUpdateIndex].Message.Text
        processMessages(currentUser.Id, chatId, lastMessageText, token)
        lastMessageId = messageId
      }
    }
  }
}
