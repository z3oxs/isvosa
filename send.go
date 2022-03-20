package isvosa

import (
    "fmt"
    "strings"
    "net/http"
    "encoding/json"
    "log"
    "io"
    "bytes"
)

func (b *Bot) Send(method interface{}) {
    endpoint := strings.Split(fmt.Sprintf("%T", method), ".")[1]
    url := fmt.Sprintf("%s/bot%s/%s", baseURL, b.Token, endpoint)

    data, _ := json.Marshal(method)
    r, err := http.Post(url, "application/json", bytes.NewBuffer(data))
    if err != nil {
        log.Fatal(err)

    }

    defer r.Body.Close()

    if r.StatusCode != 200 {
        var requestError Error
        bytes, _ := io.ReadAll(r.Body)
        json.Unmarshal(bytes, &requestError)

        log.Printf("%s [%d]", requestError.Description, requestError.ErrorCode)
    }
}

func (b *Bot) SendMessage(chatID int, text string) {
    b.Send(SendMessage { ChatID: chatID, Text: text})
}

func (b *Bot) EditMessage(chatID, messageID int, text string) {
    b.Send(EditMessageText { ChatID: chatID, MessageID: messageID, Text: text })
}

func (b *Bot) Delete(chatID, messageID int) {
    b.Send(DeleteMessage { ChatID: chatID, MessageID: messageID })
}
