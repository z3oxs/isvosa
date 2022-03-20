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

    if r.StatusCode == 404 { log.Fatal("Invalid token") }

    bytes, _ := io.ReadAll(r.Body)
    json.Unmarshal(bytes, &me)

    return me
}

func (b *Bot) GetUpdates() (Update, bool) {
    var previous Previous
    var updates Updates
    var update Update

    r, err := http.Get(fmt.Sprintf("%s/bot%s/getUpdates?offset=-1&timeout=100", baseURL, b.Token))
    if err != nil {
        log.Fatal(err)

    }
    
    if r.StatusCode == 404 { log.Fatal("Invalid token") }
    
    file, err := os.OpenFile("config.json", os.O_APPEND | os.O_CREATE, 0644)
    if err != nil {
        log.Fatal("Failed to open config.json")

    }

    defer file.Close()
    defer r.Body.Close()

    bytes, _ := io.ReadAll(r.Body)
    json.Unmarshal(bytes, &updates)
    update = updates.Update[len(updates.Update) - 1]
   
    bytes, _ = io.ReadAll(file)
    json.Unmarshal(bytes, &previous)

    if previous.Previous != update.ID {
        previous.Previous = update.ID
        
        if string(update.Message.Text) != "" && string(update.Message.Text[0]) == "/" {
            update.Command = strings.Split(update.Message.Text, " ")[0][1:]
            update.Args = strings.Split(update.Message.Text, " ")[1:]
            newFile, _ := json.Marshal(previous)
            ioutil.WriteFile("config.json", newFile, 0644)

            return update, true
        }
    }
    
    return update, false
}
