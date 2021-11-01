package tests

import ("bot/bot"
        "testing"
    	  "fmt")

var okCodes = []string{"200", "201", "202", "203", "204", "205", "206"}

func TestContains(t *testing.T) {
  okStatusCode := "200"
  badStatusCode := "400"

  if !bot.Contains(okCodes, okStatusCode){
      errorMessage := fmt.Sprintf("Expected True with statusCode %v, got False", okStatusCode)
      t.Error(errorMessage)
    }
  if bot.Contains(okCodes, badStatusCode){
      errorMessage := fmt.Sprintf("Expected False with statusCode %v, got True", badStatusCode)
      t.Error(errorMessage)
    }
}