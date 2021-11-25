package bot

import "fmt"

var howItWork = "Hello, I'am bot, what help your learn english. How it work?  You just send me word, what you wanna to learn - apple, shark, orbita, concureency, etc. And i send you as a answer sentences, what contain this word, for better understanding of context. "
var start = "Send /start to start work with bot and show list of available commands \n"
var changeLevel = "send /change_level to manualy update your englis level \n" 
var startLearning  ="send /start_learning to start education!  \n" 
var myLevel = "send /my_level  to check your current level  \n"
var startMessage = fmt.Sprintf(" %v \n..........\n %v %v %v %v", howItWork, start, startLearning, changeLevel, myLevel)
var updateEnglishLevelMessage = "Time to set your english level. Please, send me one of this nubers:\n 0 - Beginner \n 1 - Middle \n 2 - Intermediate"

