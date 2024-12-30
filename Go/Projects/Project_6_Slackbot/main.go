package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/shomali11/slacker"
	"strconv"
)

func printCommandEvents(commandEvents chan *slacker.CommandEvent){
	for e := range AnaliticsChannel{
		fmt.Println("Command events: ")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Prinln()
	}
} 

func main(){
	os.Setenv("SLACK_BOT_TOKEN","xoxb-8254882734864-8232158127075-KZkzSxEM8Qi8L2pJZTjbVntu")
	os.Setnv("SLACK_APP_TOKEN","xapp-1-A086WJZE3AQ-8232130050242-4f36053e01b9bbc8683dc15a4d4e2f0adeb93e7475d65629b06893b56c2c4259")

	bot := slacker.NewClient(os.Getnv("SLACK_BOT_TOKEN"),os.Getnv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.comand("I was born in  <year> ! ", &slacker.CommandDefinition){
		Description: "Returns the age of the user",
		Example: "I was born in 1990!",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			age = 2024 - yob
			r := fmt.Sprintf("You are %d years old!", age)
			response.Reply(r)
		}
	}
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
