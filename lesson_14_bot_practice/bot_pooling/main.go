package main

import (
	"bot_pooling/bot"
)

func main() {
	//bot.Pooling(bot.BotToken)
	//bot.getUpdates(bot.BotToken)
	//fmt.Println(bot.GetUpdates(bot.BotToken))
	bot.SendMessage("1221056d90", "this is testmessage", bot.BotToken)
}
