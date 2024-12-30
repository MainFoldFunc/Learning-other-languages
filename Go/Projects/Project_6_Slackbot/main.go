package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "github.com/shomali11/slacker"
    "strconv"
)

func printCommandEvents(commandEvents <-chan *slacker.CommandEvent) {
    for event := range commandEvents {
        fmt.Println("Command events: ")
        fmt.Println(event.Timestamp)
        fmt.Println(event.Command)
        fmt.Println(event.Parameters)
        fmt.Println(event.Event)
        fmt.Println() // Fixed typo from "Println" to "println"
    }
}

func main() {
    // Set environment variables (note: this should not be hardcoded for security reasons)
    os.Setenv("SLACK_BOT_TOKEN", "xoxb-8254882734864-8232158127075-KZkzSxEM8Qi8L2pJZTjbVntu")
    os.Setenv("SLACK_APP_TOKEN", "xapp-1-A086WJZE3AQ-8232130050242-4f36053e01b9bbc8683dc15a4d4e2f0adeb93e7475d65629b06893b56c2c4259")

    // Initialize the bot client
    bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

    // Start printing command events in a goroutine
    go printCommandEvents(bot.CommandEvents())

    // Define a command for calculating age
    bot.Command("I was born in <year>!", &slacker.CommandDefinition{
        Description: "Returns the age of the user",
        Examples:    []string{"I was born in 1990!"}, // Corrected field name
        Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
            year := request.Param("year")
            yob, err := strconv.Atoi(year)
            if err != nil {
                fmt.Println("Error: ", err)
                response.Reply("Sorry, I couldn't process your birth year.")
                return
            }
            age := 2024 - yob // Declare and calculate age
            r := fmt.Sprintf("You are %d years old!", age)
            response.Reply(r)
        },
    })

    // Listen for commands
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    err := bot.Listen(ctx)
    if err != nil {
        log.Fatal(err)
    }
}
