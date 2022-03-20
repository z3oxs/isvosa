<div align="center">
    <img width="500" src="isvosa.png" />
    <h3>A performatic library to develop Telegram bots.</h3>
</div>

<br><br>
## Install
```go
go get -u github.com/z3oxs/isvosa
```

<br><br>
## 📃 Documentation
<a id="summary" /><br>
- [Getting started](#getting-started)
- [Bot information](#bot-information)
- [Sending messages and other medias](#sending)
- [Editting](#editting)
- [Other actions](#other-actions)

<br><br>
<a id="getting-started" />
### Getting started
Initiate a new project and import the package
```go
package main

// Make sure you getted the package before getting crazy about that
import "github.com/z3oxs/isvosa"

func main() {
    // Define the token inside the bot struct
    bot := isvosa.Bot {
        Token: "YOUR TOKEN HERE"
    }
    
    // Simple updates handler to get information
    for {
        // 'Update' var will receive all updates information and 'newUpdate' will receive a boolean confirming if
        // the new update is really a new update
        update, newUpdate := bot.GetUpdates()
        
        // Checking if is a trully new update
        if newUpdate {
            // Handling all valid commands entries (Contains '/' on first character)
            switch update.Command {
                // Will send a simple message
                case "ping":
                    bot.Send(isvosa.SendMessage {
                        ChatID: update.Message.Chat.ID,
                        Text: "test",
                    })

                // Sending a message replying will first element of the arguments, that is parsed automatically
                // for all valid commands
                case "echo":
                    bot.Send(isvosa.SendMessage {
                        ChatID: update.Message.Chat.ID,
                        Text: update.Args[0],
                    })
        }
    }
}
```
<br><br>
<a id="bot-information" />
[Back to summary](#summary)
### Bot information
Getting information about the user of the bot
```go
func me(bot isvosa.Bot) {
    // Getting information
    info := bot.GetMe()
    
    // Will print ID from the bot
    fmt.Println(info.Me.ID)
    
    // Will print first name from the bot
    fmt.Println(info.Me.FirstName)
    
    // More information: https://core.telegram.org/bots/api#user
}
```
<br><br>
<a id="sending" />
[Back to summary](#summary)
### Sending messages and other medias
Basic about sending message and media with your bot
```go
// To send complete messages, with complementar methods, more: https://core.telegram.org/bots/api#sendmessage
bot.Send(isvosa.SendMessage {
    ChatID: update.Message.Chat.ID,
    Text: "Text",
})

// To send photos, more: https://core.telegram.org/bots/api#sendphoto
bot.Send(isvosa.SendPhoto {
    ChatID: update.Message.Chat.ID,
    Photo: "A URL or a InputMedia (https://core.telegram.org/bots/api#inputmedia)",
})
// To send a inline keyboard, more: https://core.telegram.org/bots/api#inlinekeyboardmarkup
bot.Send(isvosa.SendMessage {
    ChatID: update.Message.Chat.ID,
    Text: "Inline keyboard test",
    ReplyMarkup: isvosa.InlineKeyboardMarkup {
        InlineKeyboard: [][]isvosa.InlineKeyboardButton {
            // First row, simple URL button
            []isvosa.InlineKeyboardButton {
                { Text: "Row 1", URL: "https://test.com" },
            },
            // Second row, with 2 lines (You can define any size for your inline keyboard, just
            // add '[]isvosa.InlineKeyboardButton' to insert)
            []isvosa.InlineKeyboardButton {
                { Text: "Row 2", URL: "https://test.com" },
                { Text: "Row 2 Line 2", URL: "https://test.com" }
            },
        },
    },
})
```

#### All available formats:
- isvosa.SendMessage -> https://core.telegram.org/bots/api#sendmessage
- isvosa.SendPhoto -> https://core.telegram.org/bots/api#sendphoto
- isvosa.SendAudio -> https://core.telegram.org/bots/api#sendaudio
- isvosa.SendDocument -> https://core.telegram.org/bots/api#senddocument
- isvosa.SendVideo -> https://core.telegram.org/bots/api#sendvideo
- isvosa.SendAnimation -> https://core.telegram.org/bots/api#sendanimation
- isvosa.SendVoice -> https://core.telegram.org/bots/api#sendvoice
- isvosa.SendVideoNote -> https://core.telegram.org/bots/api#sendvideonote
- isvosa.SendMediaGroup -> https://core.telegram.org/bots/api#sendmediagroup
- isvosa.SendLocation -> https://core.telegram.org/bots/api#sendlocation
- isvosa.SendVenue -> https://core.telegram.org/bots/api#sendvenue
- isvosa.SendContact -> https://core.telegram.org/bots/api#sendcontact
- isvosa.SendPoll -> https://core.telegram.org/bots/api#sendpoll
- isvosa.SendDice -> https://core.telegram.org/bots/api#senddice
- isvosa.SendChatAction -> https://core.telegram.org/bots/api#sendchataction

<br><br>
<a id="editting" />
[Back to summary](#summary)
### Editting
Editing existing messages information
```go
// Will change the content of next message (based on your input)
bot.Send(isvosa.EditMessageText {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
    Text: "New text",
})

// To change any media
bot.Send(isvosa.EditMessageMedia {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
    Media: "A URL or InputMedia(https://core.telegram.org/bots/api#inputmedia)",
})
```

#### All available formats:
- isvosa.EditMessageText -> https://core.telegram.org/bots/api#editmessagetext
- isvosa.EditMessageCaption -> https://core.telegram.org/bots/api#editmessagecaption
- isvosa.EditMessageMedia -> https://core.telegram.org/bots/api#editmessagemedia
- isvosa.EditMessageReplyMarkup -> https://core.telegram.org/bots/api#editmessagereplymarkup

<br><br>
<a id="other-actions" />
[Back to summary](#summary)
### Other actions
Some other actions you can do with your bot
```go
// Will stop any poll
bot.Send(isvosa.StopPoll {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
})

// Will delete a message
bot.Send(isvosa.DeleteMessage {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
})
```

**New features and information coming soon.**
