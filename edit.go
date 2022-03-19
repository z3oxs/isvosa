package isvosa

import (
    "fmt"
    "log"
)

func (b *Bot) Edit(Struct interface{}) {
    endpoint := ""
    switch fmt.Sprintf("%T", Struct) {
        case "isvosa.EditMessageText": endpoint = "editMessageText"
        case "isvosa.EditMessageCaption": endpoint = "editMessageCaption"
        case "isvosa.EditMessageMedia": endpoint = "editMessageMedia"
        case "isvosa.EditMessageReplyMarkup": endpoint = "editMessageReplyMarkup"
        default: log.Fatal("Invalid struct parsed to \"Edit\"")
    }

    request(Struct, b.Token, endpoint)
}

func (b *Bot) EditMessage(chatID int, messageID int, text string) {
    request(EditMessageText{ ChatID: chatID, MessageID: messageID, Text: text }, b.Token, "editMessageText")
}
