package bot

const rootURL = "https://api.telegram.org/bot"
var BotToken = "2125376282:AAFpG_uW9d0qBoRFOT0cbEaogM2m8tiFblU"

var englishSentences = sentencesFromJson("sentences.json")

var okCodes = []string{"200", "201", "202", "203", "204", "205", "206"}
var botCommands = []string{"/start", "/change_level", "/start_learning", "/my_level"}


type Sentence struct {Text string
                      Level int}

type User struct{
          Id int             `json:"id"`
          Username string    `json:"username"`
          Level int          `json:"englishLevel"`
          // UpdateInProgress bool
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