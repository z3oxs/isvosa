package isvosa

import (
    "net/http"
    "fmt"
    "bytes"
    "encoding/json"
    "log"
)

func request(anyStruct interface{}, endpoint, token string) {
    data, _ := json.Marshal(anyStruct)

    r, err := http.Post(fmt.Sprintf("%s/bot%s/%s", baseURL, token, endpoint), "application/json", bytes.NewBuffer(data))
    if err != nil {
        log.Fatal(err)

    }

    defer r.Body.Close()

    if r.StatusCode != 200 {
        log.Printf("Invalid call/body (Status: %d / Endpoint: %s)", r.StatusCode, endpoint)

    }
}

func (b *Bot) SendMessage(message MessageStruct) {
    request(message, "sendMessage", b.Token)
}

func (b *Bot) SendPhoto(photo PhotoStruct) {
    request(photo, "sendPhoto", b.Token)
}

func (b *Bot) SendAudio(audio AudioStruct) {
    request(audio, "sendAudio", b.Token)
}

func (b *Bot) SendDocument(document DocumentStruct) {
    request(document, "sendDocument", b.Token)
}

func (b *Bot) SendVideo(video VideoStruct) {
    request(video, "sendVideo", b.Token)
}
