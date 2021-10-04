package bot

const rootURL = "https://api.telegram.org/bot"

var BotToken = "2023508511:AAGAs6aPfN3GJXGJFEmK6bjUtmJCgJxCYvw"

var okCodes = []string{"200", "201", "202", "203", "204", "205", "206"}

type Chat struct {
	Id int `json:"id"`
}

type Message struct {
	Id   int    `json:"message_id"`
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Updates struct {
	Updates []Update `json:"result"`
}
