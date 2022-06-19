<div align="center">
    <h3>A simple library to fast develop Telegram bots.</h3>
    <a href="https://pkg.go.dev/github.com/z3oxs/isvosa">
        <img src="https://pkg.go.dev/badge/github.com/z3oxs/isvosa.svg" />
    </a>
    <a href="https://www.codefactor.io/repository/github/z3oxs/isvosa">
        <img src="https://www.codefactor.io/repository/github/z3oxs/isvosa/badge" alt="CodeFactor" />
    </a>
</div>

<br><br>
## Install
```go
go get -u github.com/z3oxs/isvosa
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
import "github.com/z3oxs/isvosa"

func main() {
    // Define the token inside the bot struct
    bot := isvosa.Bot {
        Token: "YOUR TOKEN HERE"
    }
    
    // Is required 2 parameters, the first is the command, if anyone send "/ping" to the bot, will be handled
    // based on the second parameter, the function, requiring 3 parameters, can be with any name, but is
    // required to be 3 and with types "*isvosa.Bot", "*isvosa.Message" and "[]string", respectively
    bot.Add("ping", func(bot *isvosa.Bot, msg *isvosa.Message, args []string) {
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
    "github.com/z3oxs/isvosa"
)

func main() {
    // Initialize the variable with a Bot Type
    bot := isvosa.Bot {
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
    "github.com/z3oxs/isvosa"
)

// Can be empty or doing any stuff, we will list all available commands, only needing to be defined
func Setup() {
    for _, f := range isvosa.Command() {
        fmt.Printf("Loaded %v\n", f.Command)
    }
}
```

<br><br>
commands/ping.go: (a available command that will reply with "pong")
```go
package commands

import "github.com/z3oxs/isvosa"

// This function will run in the same moment that any command from the package was been summoned,
// we will use to add the command to the handler
func init() {
    isvosa.Add(isvosa.Command {
        // If "/ping" was been sended to bot, he will handle the request as a valid command
        Command: "ping",
        // You can pass a anonymous function or a existing function, only requiring receive
        // 3 parameters, bot as *isvosa.Bot, msg as *isvosa.Message and args as []string, the 
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
func ping(bot *isvosa.Bot, msg *isvosa.Message, args []string) {
    bot.SendMessage(msg.Chat.ID, "pong!")
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
<a id="mod" />
[Back to summary](#mod)
### Moderation
Moderation binds for take control of your chat
```go
// A function to ensure if the command sender is an admin
func ensureAdmin(bot *isvosa.Bot, msg *isvosa.Message, adminID int) bool {
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
bot.Add("ban", func(bot *isvosa.Bot, msg *isvosa.Message, args []string) {
    admin := ensureAdmin(bot, msg, msg.From.ID)

    if admin {
        // You have 3 methods to send a valid ID to ban, replying to user message, parsing his ID or
        // logging all messages on a database with all user information (after had sended a message)
        // and searching for an ID on database
        bot.Ban(msg.Chat.ID, msg.ReplyToMessage.From.ID)

    }
}

// A simple "unban" command
bot.Add("ban", func(bot *isvosa.Bot, msg *isvosa.Message, args []string) {
    admin := ensureAdmin(msg.From.ID)

    if admin {
        bot.Unban(msg.Chat.ID, msg.ReplyToMessage.From.ID)

    }
}

// A detailed "ban" command, all moderation structs are sended using "bot.Send"
bot.Add("ban", func(bot *isvosa.Bot, msg *isvosa.Message, args []string) {
    admin := ensureAdmin(msg.From.ID)
    
    if admin {
        // Parsing a struct using "Send" method
        bot.Send(isvosa.BanChatMember {
            ChatID: msg.Chat.ID,
            UserID: msg.ReplyToMessage.From.ID,
        })
    }
})

// A detailed "unban" command
bot.Add("unban", func(bot *isvosa.Bot, msg *isvosa.Message, args []string) {
    admin := ensureAdmin(msg.From.ID)
    
    if admin {
        bot.Send(isvosa.UnbanChatMember {
            ChatID: msg.Chat.ID,
            UserID: msg.ReplyToMessage.From.ID,
        })
    }
})
```

#### All available actions:
- isvosa.BanChatMember -> https://core.telegram.org/bots/api#banchatmember
- isvosa.UnbanChatMember -> https://core.telegram.org/bots/api#unbanchatmember
- isvosa.RestrictChatMember -> https://core.telegram.org/bots/api#restrictchatmember
- isvosa.PromoteChatMember -> https://core.telegram.org/bots/api#promotechatmember
- isvosa.SetChatAdministratorCustomTitle -> https://core.telegram.org/bots/api#setchatadministratorcustomtitle
- isvosa.BanChatSenderChat -> https://core.telegram.org/bots/api#banchatsenderchat
- isvosa.UnbanChatSenderChat -> https://core.telegram.org/bots/api#unbanchatsenderchat
- isvosa.SetChatPermissions -> https://core.telegram.org/bots/api#setchatpermissions
- isvosa.DeleteChatPhoto -> https://core.telegram.org/bots/api#deletechatphoto
- isvosa.SetChatTitle -> https://core.telegram.org/bots/api#setchattitle
- isvosa.SetChatDescription -> https://core.telegram.org/bots/api#setchatdescription
- isvosa.PinChatMessage -> https://core.telegram.org/bots/api#pinchatmessage
- isvosa.UnpinChatMessage -> https://core.telegram.org/bots/api#unpinchatmessage
- isvosa.UnpinAllChatMessages -> https://core.telegram.org/bots/api#unpinallchatmessages
- isvosa.SetChatStickerSet -> https://core.telegram.org/bots/api#setchatstickerset
- isvosa.DeleteChatStickerSet -> https://core.telegram.org/bots/api#deletechatstickerset
- isvosa.SetMyCommands -> https://core.telegram.org/bots/api#setmycommands
- isvosa.DeleteMyCommands -> https://core.telegram.org/bots/api#deletemycommands
- isvosa.LeaveChat -> https://core.telegram.org/bots/api#leavechat

<br><br>
<a id="editting" />
[Back to summary](#summary)
### Editting
Editing existing messages information
```go
// Simple message edit
bot.EditMessage(update.Message.Chat.ID, update.Message.ID, "New content!")

// Can send a completely change to the content of messages (based on your input), more: https://core.telegram.org/bots/api#editmessagetext
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
<a id="get" />
[Back to summary](#summary)
### Getting information
Functions to get usual information, or not, i don't know :/
```go
// Will fetch all parsed user ID profile photos
photos := bot.GetUserProfilePhotos(isvosa.GetUserProfilePhotos {
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
bot.Send(isvosa.StopPoll {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
})

// Will simple delete a message
bot.Delete(update.Message.Chat.ID, update.Message.ID)

// Will delete a message with more parameters, more: https://core.telegram.org/bots/api#deletemessage
bot.Send(isvosa.DeleteMessage {
    ChatID: update.Message.Chat.ID,
    MessageID: update.Message.ID,
})
```
