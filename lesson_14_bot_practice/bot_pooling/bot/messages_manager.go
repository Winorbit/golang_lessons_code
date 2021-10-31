package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func contains(elements []string, element string) bool {
	for _, el := range elements {
		if el == element {
			return true
		}
	}
	return false
}

func GetUpdates(token string) Updates {
	updatesUrl := fmt.Sprintf("%s%s/%s", rootURL, token, "getUpdates")
	resp, _ := http.Get(updatesUrl)

	var updates Updates
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&updates)
	return updates
}

func SendMessage(chatId string, message string, token string) {
	sendMessageURL := fmt.Sprintf("%s%s/sendMessage", rootURL, token)
	res, err := http.PostForm(sendMessageURL,
		url.Values{"chat_id": {chatId},
			"text": {message}})

	if err != nil {
		fmt.Println(err.Error())
	}

	statusCode := res.Status[0:3]

	codeOk := contains(okCodes, statusCode)

	if codeOk != true {
		fmt.Printf("Request failed with statuscode %s", statusCode)
	}

}

// func processMessages(chatId string, message string, token string){
//   if message == "Hello"{
//     sendMessage(chatId, "Your welcome!", token)
//   } else {
//     sendMessage(chatId, message, token)
//   }
// }

// func Pooling(token string){
//   ...
// }
