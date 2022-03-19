// A performatic Telegram API to build bots
package isvosa

import (
    "net/http"
    "fmt"
    "encoding/json"
    "io"
    "log"
    "strings"
    "os"
    "io/ioutil"
)

func (b *Bot) GetMe() Me {
    var me Me

    r, err := http.Get(fmt.Sprintf("%s/bot%s/getMe", baseURL, b.Token))
    if err != nil {
        log.Fatal(err)
    }

    defer r.Body.Close()

    if r.StatusCode == 404 {
        log.Fatal("Invalid token")

    }

    bytes, _ := io.ReadAll(r.Body)
    err = json.Unmarshal(bytes, &me)
    if err != nil {
        log.Fatal(err)

    }

    return me
}

func (b *Bot) GetUpdates() (Update, bool) {
    var updates Updates
    var update Update

    r, err := http.Get(fmt.Sprintf("%s/bot%s/getUpdates?timeout=100", baseURL, b.Token))
    if err != nil {
        log.Fatal(err)

    }

    defer r.Body.Close()
    
    if r.StatusCode == 404 {
        log.Fatal("Invalid token")

    }

    bytes, _ := io.ReadAll(r.Body)
    err = json.Unmarshal(bytes, &updates)
    if err != nil {
        log.Fatal(err)

    }

    update = updates.Update[len(updates.Update) - 1]

    var previous Previous
    file, err := os.Create("config.json")
    if err != nil {
        log.Fatal("Failed to open config.json")

    }
    
    bytes, _ = io.ReadAll(file)
    json.Unmarshal(bytes, &previous)

    if previous.Previous != update.ID && string(update.Message.Text) != "" && string(update.Message.Text[0]) == "/" {
        update.Command = strings.Split(update.Message.Text, " ")[0][1:]
        update.Args = strings.Split(update.Message.Text, " ")[1:]
        previous.Previous = update.ID
        jsonFile, _ := json.Marshal(previous)
        ioutil.WriteFile("config.json", jsonFile, 0644)

        return update, true
    }
    
    return update, false
}
