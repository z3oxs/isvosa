package isvosa

import (
    "net/http"
    "log"
    "fmt"
    "encoding/json"
    "io"
)

// Get all bot informations when the token is valid
func (b *Bot) GetMe() Me {
    var me Me
    r, err := http.Get(fmt.Sprintf("%s/bot%s/getMe", baseURL, b.Token))
    if err != nil {
        log.Fatal(err)
    }

    defer r.Body.Close()

    if r.StatusCode == 404 { log.Fatal("Invalid token") }

    bytes, _ := io.ReadAll(r.Body)
    json.Unmarshal(bytes, &me)

    return me
}
