package bot

import ("log"
        "os")

const rootURL = "https://api.telegram.org/bot"
var BotToken = "2125376282:AAFpG_uW9d0qBoRFOT0cbEaogM2m8tiFblU"

var sessions = map[int]*User{}
var englishSentences = sentencesFromJson("sentences.json")

var okCodes = []string{"200", "201", "202", "203", "204", "205", "206"}
var botCommands = []string{"/start", "/change_level", "/start_learning", "/my_level"}


type Sesssion struct{SessionID int
                     UserInfo User}

type Sentence struct {Text string
                      Level int}

var availableLevels = []int{0, 1, 2}

type User struct{
          Id int             `json:"id"`
          IsBot bool         `json:"is_bot"`
          Username string    `json:"username"`
          Level int          `json:"englishLevel"`
          UpdateInProgress bool
        }

type Chat struct {
  Id int `json:"id"`
}

type Message struct {
  Id    int     `json:"message_id"`
  Text  string  `json:"text"`
  Chat  Chat    `json:"chat"`
  User  User    `json:"from"`
}

type Update struct {
  UpdateId int     `json:"update_id"`
  Message  Message `json:"message"`
}

type Updates struct {
  Updates []Update  `json:"result"`
}


var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)