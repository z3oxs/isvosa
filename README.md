<div align="center">
    <h3>A simple library to fast develop Telegram bots.</h3>
    <a href="https://pkg.go.dev/github.com/z3oxs/telego">
        <img src="https://pkg.go.dev/badge/github.com/z3oxs/telego.svg" />
    </a>
    <a href="https://www.codefactor.io/repository/github/z3oxs/telego">
        <img src="https://www.codefactor.io/repository/github/z3oxs/telego/badge" alt="CodeFactor" />
    </a>
</div>

<br><br>
## Install
```go
go get -u github.com/z3oxs/telego
```

<br><br>
## ♻️ Changelog v0.1.7
**For some reason golang cache all the updates, on commands like ban, unban you need
to send another message to overwrite the last update, or the commands will be looped.
Only affecting commands that send a notification of some modification on settings or
member restriction**

- Moderation binds [more](#mod)
    - banChatMember
    - unbanChatMember
    - restrictChatMember
    - promoteChatMember
    - setChatAdministratorCustomTitle
    - banChatSenderChat
    - unbanChatSenderChat
    - setChatPermissions
    - deleteChatPhoto
    - setChatTitle
    - setChatDescription
    - pinChatMessage
    - unpinChatMessage
    - unpinAllChatMessages
    - setChatStickerSet
    - deleteChatStickerSet
    - setMyCommands
    - deleteMyCommands
    - leaveChat

<br><br>
## 📃 Documentation
<a id="summary" /><br>
- [Getting started](#getting-started)
- [Modular handler](#modular-handler)
- [Bot information](#bot-information)
- [Sending messages and other medias](#sending)
- [Moderation](#mod)
- [Editting](#editting)
- [Getting information](#get)
- [Other actions](#other-actions)

<br><br>
<a id="getting-started" />
### Getting started
Initiate a new project and import the package
```go
package main

// Make sure you getted the package before getting crazy about that
import "github.com/z3oxs/telego"

func main() {
    // Define the token inside the bot struct
    bot := telego.Bot {
        Token: "YOUR TOKEN HERE"
    }
    
    // Is required 2 parameters, the first is the command, if anyone send "/ping" to the bot, will be handled
    // based on the second parameter, the function, requiring 3 parameters, can be with any name, but is
    // required to be 3 and with types "*telego.Bot", "*telego.Message" and "[]string", respectively
    bot.Add("ping", func(bot *telego.Bot, msg *telego.Message, args []string) {
        bot.SendMessage(msg.Chat.ID, "pong!")
    })
    
    // Will nonstop polling, automatic handling all commands and messages received
    bot.Start()
}
```

<br><br>
<a id="modular-handler" />
[Back to summary](#summary)
### Modular handler
Handling commands with functions outside main.go **All command files need to be of the same package and in the same directory**

main.go:
```go
package main

import (
    // A module with any name you choose, i'll call "commands"
    "example/commands"
    "github.com/z3oxs/telego"
)

func main() {
    // Initialize the variable with a Bot Type
    bot := telego.Bot {
        Token: "YOUR TOKEN HERE" 
    }

    // Create a exportable function from your module, you can choose any name for the method, but
    // is **required**, without a function to call the module, the commands will not be exported
    commands.Setup()

    // Nonstop polling to automatic get new and handle updates
    bot.Start()
}
```

<br><br>
commands/commands.go: (a file inside another path inside the project)
```go
package commands

import (
    "fmt"
    "github.com/z3oxs/telego"
)

// Can be empty or doing any stuff, we will list all available commands, only needing to be defined
func Setup() {
    for _, f := range telego.Command() {
        fmt.Printf("Loaded %v\n", f.Command)
    }
}
```

<br><br>
commands/ping.go: (a available command that will reply with "pong")
```go
package commands

import "github.com/z3oxs/telego"

// This function will run in the same moment that any command from the package was been summoned,
// we will use to add the command to the handler
func init() {
    telego.Add(telego.Command {
        // If "/ping" was been sended to bot, he will handle the request as a valid command
        Command: "ping",
        // You can pass a anonymous function or a existing function, only requiring receive
        // 3 parameters, bot as *telego.Bot, msg as *telego.Message and args as []string, the 
        // parameter name can be anyone, only requiring 3 parameters with these 3 types, 
        // respectively
        Run: ping,
        // The description is fully optional, doens't handled with the function, only for context
        // purposes
        Description: "Reply a text contening 'pong!'"
    })
}

// If you choose creating a non-anonymous function, that is the format you need to use to
// handle valid commands
func ping(bot *telego.Bot, msg *telego.Message, args []string) {
    bot.SendMessage(msg.Chat.ID, "pong!")
}
```

<br><br>
<a id="bot-information" />
[Back to summary](#summary)
### Bot information
Getting information about the user of the bot
```go
func me(bot telego.Bot) {
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
// To send photos, more: https://core.telegram.org/bots/api#sendphoto
bot.Send(telego.SendPhoto {
    ChatID: update.Message.Chat.ID,
    Photo: "A URL or a InputMedia (https://core.telegram.org/bots/api#inputmedia)",
})

// To send a inline keyboard, more: https://core.telegram.org/bots/api#inlinekeyboardmarkup
bot.Send(telego.SendMessage {
    ChatID: update.Message.Chat.ID,
    Text: "Inline keyboard test",
    ReplyMarkup: telego.InlineKeyboardMarkup {
        InlineKeyboard: [][]telego.InlineKeyboardButton {
            // First row, simple URL button
            []telego.InlineKeyboardButton {
                { Text: "Row 1", URL: "https://test.com" },
            },
            // Second row, with 2 lines (You can define any size for your inline keyboard, just
            // add '[]telego.InlineKeyboardButton' to insert)
            []telego.InlineKeyboardButton {
                { Text: "Row 2", URL: "https://test.com" },
                { Text: "Row 2 Line 2", URL: "https://test.com" }
            },
        },
    },
})
```

#### All available formats:
- telego.SendMessage -> https://core.telegram.org/bots/api#sendmessage
- .SendPhoto -> https://core.telegram.org/bots/api#sendphoto
- telego.SendAudio -> https://core.telegram.org/bots/api#sendaudio
- telego.SendDocument -> https://core.telegram.org/bots/api#senddocument
- telego.SendVideo -> https://core.telegram.org/bots/api#sendvideo
- telego.SendAnimation -> https://core.telegram.org/bots/api#sendanimation
- telego.SendVoice -> https://core.telegram.org/bots/api#sendvoice
- telego.SendVideoNote -> https://core.telegram.org/bots/api#sendvideonote
- telego.SendMediaGroup -> https://core.telegram.org/bots/api#sendmediagroup
- telego.SendLocation -> https://core.telegram.org/bots/api#sendlocation
- telego.SendVenue -> https://core.telegram.org/bots/api#sendvenue
- telego.SendContact -> https://core.telegram.org/bots/api#sendcontact
- telego.SendPoll -> https://core.telegram.org/bots/api#sendpoll
- telego.SendDice -> https://core.telegram.org/bots/api#senddice
- telego.SendChatAction -> https://core.telegram.org/bots/api#sendchataction

<br><br>
<a id="mod" />
[Back to summary](#mod)
### Moderation
Moderation binds for take control of your chat
```go
// A function to ensure if the command sender is an admin
func ensureAdmin(bot *telego.Bot, msg *telego.Message, adminID int) bool {
    // Getting a array returning all admins
    admins := bot.GetChatAdministrators(msg.Chat.ID)
    
    // Checking if the user that sended the command is admin
    for _, i := range admins {
        if msg.From.ID == adminID {
            return true

        }
    }

    return false
}

// A simple "ban" command
bot.Add("ban", func(bot *telego.Bot, msg *telego.Message, args []string) {
    admin := ensureAdmin(bot, msg, msg.From.ID)

    if admin {
        // You have 3 methods to send a valid ID to ban, replying to user message, parsing his ID or
        // logging all messages on a database with all user information (after had sended a message)
        // and searching for an ID on database
        bot.Ban(msg.Chat.ID, msg.ReplyToMessage.From.ID)

    }
}

// A simple "unban" command
bot.Add("ban", func(bot *telego.Bot, msg *telego.Message, args []string) {
    admin := ensureAdmin(msg.From.ID)

    if admin {
        bot.Unban(msg.Chat.ID, msg.ReplyToMessage.From.ID)

    }
}

// A detailed "ban" command, all moderation structs are sended using "bot.Send"
bot.Add("ban", func(bot *telego.Bot, msg *telego.Message, args []string) {
    admin := ensureAdmin(msg.From.ID)
    
    if admin {
        // Parsing a struct using "Send" method
        bot.Send(telego.BanChatMember {
            ChatID: msg.Chat.ID,
            UserID: msg.ReplyToMessage.From.ID,
        })
    }
})

// A detailed "unban" command
bot.Add("unban", func(bot *telego.Bot, msg *telego.Message, args []string) {
    admin := ensureAdmin(msg.From.ID)
    
    if admin {
        bot.Send(telego.UnbanChatMember {
            ChatID: msg.Chat.ID,
            UserID: msg.ReplyToMessage.From.ID,
        })
    }
})
```

#### All available actions:
- telego.BanChatMember -> https://core.telegram.org/bots/api#banchatmember
- telego.UnbanChatMember -> https://core.telegram.org/bots/api#unbanchatmember
- telego.RestrictChatMember -> https://core.telegram.org/bots/api#restrictchatmember
- telego.PromoteChatMember -> https://core.telegram.org/bots/api#promotechatmember
- telego.SetChatAdministratorCustomTitle -> https://core.telegram.org/bots/api#setchatadministratorcustomtitle
- telego.BanChatSenderChat -> https://core.telegram.org/bots/api#banchatsenderchat
- telego.UnbanChatSenderChat -> https://core.telegram.org/bots/api#unbanchatsenderchat
- telego.SetChatPermissions -> https://core.telegram.org/bots/api#setchatpermissions
- telego.DeleteChatPhoto -> https://core.telegram.org/bots/api#deletechatphoto
- telego.SetChatTitle -> https://core.telegram.org/bots/api#setchattitle
- telego.SetChatDescription -> https://core.telegram.org/bots/api#setchatdescription
- telego.PinChatMessage -> https://core.telegram.org/bots/api#pinchatmessage
- telego.UnpinChatMessage -> https://core.telegram.org/bots/api#unpinchatmessage
- telego.UnpinAllChatMessages -> https://core.telegram.org/bots/api#unpinallchatmessages
- telego.SetChatStickerSet -> https://core.telegram.org/bots/api#setchatstickerset
- telego.DeleteChatStickerSet -> https://core.telegram.org/bots/api#deletechatstickerset
- telego.SetMyCommands -> https://core.telegram.org/bots/api#setmycommands
- telego.DeleteMyCommands -> https://core.telegram.org/bots/api#deletemycommands
- telego.LeaveChat -> https://core.telegram.org/bots/api#leavechat

<br><br>
<a id="editting" />
[Back to summary](#summary)
### Editting
Editing existing messages information
```go
// Simple message edit
bot.EditMessage(update.Message.Chat.ID, update.Message.ID, "New content!")

// Can send a completely change to the content of messages (based on your input), more: https://core.telegram.org/bots/api#editmessagetext
bot.Send(telego.EditMessageText {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
    Text: "New text",
})

// To change any media
bot.Send(telego.EditMessageMedia {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
    Media: "A URL or InputMedia(https://core.telegram.org/bots/api#inputmedia)",
})
```

#### All available formats:
- telego.EditMessageText -> https://core.telegram.org/bots/api#editmessagetext
- telego.EditMessageCaption -> https://core.telegram.org/bots/api#editmessagecaption
- telegoEditMessageMedia -> https://core.telegram.org/bots/api#editmessagemedia
- telego.EditMessageReplyMarkup -> https://core.telegram.org/bots/api#editmessagereplymarkup

<br><br>
<a id="get" />
[Back to summary](#summary)
### Getting information
Functions to get usual information, or not, i don't know :/
```go
// Will fetch all parsed user ID profile photos
photos := bot.GetUserProfilePhotos(telego.GetUserProfilePhotos {
    UserID: <Some ID here>
})

// Will fetch information from the parsed chat
chat := bot.GetChat(<Chat ID>)

// Will fetch a array of all admins belong to parsed chat
admins := bot.GetChatAdministrators(<Chat ID>)

// Will return a integer that is all member count of parsed chat
count := bot.GetChatMemberCount(<Chat ID>)

// Will return a array of valid commands set by BotFather
commands := bot.GetMyCommands("scope") // Available scopes: https://core.telegram.org/bots/api#botcommandscope
```

<br><br>
<a id="other-actions" />
[Back to summary](#summary)
### Other actions
Some other actions you can do with your bot
```go
// Will stop any poll
bot.Send(telego.StopPoll {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
})

// Will simple delete a message
bot.Delete(update.Message.Chat.ID, update.Message.ID)

// Will delete a message with more parameters, more: https://core.telegram.org/bots/api#deletemessage
bot.Send(telego.DeleteMessage {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
})
```
