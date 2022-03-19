package isvosa

import (
    "fmt"
    "log"
)

func (b *Bot) Send(Struct interface{}) {
    endpoint := ""
    switch fmt.Sprintf("%T", Struct) {
        case "isvosa.SendMsg": endpoint = "sendMessage"
        case "isvosa.SendPhoto": endpoint = "sendPhoto"
        case "isvosa.SendAudio": endpoint = "sendAudio"
        case "isvosa.SendDocument": endpoint = "sendDocument"
        case "isvosa.SendVideo": endpoint = "sendVideo"
        case "isvosa.SendAnimation": endpoint = "sendAnimation"
        case "isvosa.SendVoice": endpoint = "sendVoice"
        case "isvosa.SendVideoNote": endpoint = "sendVideoNote"
        case "isvosa.SendMediaGroup": endpoint = "sendMediaGroup"
        case "isvosa.SendLocation": endpoint = "sendLocation"
        case "isvosa.SendVenue": endpoint = "sendVenue"
        case "isvosa.SendContact": endpoint = "sendContact"
        case "isvosa.SendPoll": endpoint = "sendPoll"
        case "isvosa.SendDice": endpoint = "sendDice"
        case "isvosa.SendChatAction": endpoint = "sendChatAction"
        default: log.Fatal("Invalid struct parsed to \"Send\"")
    }

    request(Struct, b.Token, endpoint)
}

func (b *Bot) SendMessage(chatID int, text string) {
    request(SendMsg{ ChatID: chatID, Text: text }, b.Token, "sendMessage")
}
