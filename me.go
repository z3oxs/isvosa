package isvosa

import (
    "net/http"
    "log"
    "fmt"
    "encoding/json"
    "io"
)

func (b *Bot) GetMe() Me {
    r, err := http.Get(fmt.Sprintf("%s/bot%s/getMe", baseURL, b.Token))
    if err != nil {
        log.Fatal(err)

    } else if r.StatusCode == 404 {
        log.Fatal("Invalid token.")

    } 

    defer r.Body.Close()

    var me Me
    bytes, _ := io.ReadAll(r.Body)
    json.Unmarshal(bytes, &me)

    return me
}
