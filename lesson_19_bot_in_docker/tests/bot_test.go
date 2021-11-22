package tests

import ("bot_pooling/bot"
        "testing"
        "log"
        "os"
    	  "fmt")

var logger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
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

  log.Println("Testing function Contains passed.")
  logger.Println("Testing function Contains passed.")
}