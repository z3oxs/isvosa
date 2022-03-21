// A simple and performatic Telegram Bot API
package isvosa

import (
    "net/http"
    "fmt"
    "encoding/json"
    "io"
    "log"
    "strings"
)

// Will handling all the updates, being the main pillar of the package doingl
// a nonstop polling based on a for loop and getting all the new updates and
// testing if is really a new update
func (b *Bot) Start() {
    for {
        var updates Updates
        var update Update

        r, err := http.Get(fmt.Sprintf("%s/bot%s/getUpdates?offset=-1&timeout=100", baseURL, b.Token))
        if err != nil {
            log.Fatal(err)

        }
        
        defer r.Body.Close()
        
        if r.StatusCode == 404 { log.Fatal("Invalid token") }

        bytes, _ := io.ReadAll(r.Body)
        json.Unmarshal(bytes, &updates)
        update = updates.Update[len(updates.Update) - 1]
        
        if previousID == 0 {
            previousID = update.ID
        
            continue

        } else if previousID != update.ID {
            previousID = update.ID
            
            if string(update.Message.Text) != "" && string(update.Message.Text[0]) == "/" {
                update.Command = strings.Split(update.Message.Text, " ")[0][1:]
                update.Args = strings.Split(update.Message.Text, " ")[1:]

                if len(update.Args) == 0 {
                    update.Args = []string{""}

                }

                for _, f := range handler.Commands {
                    if update.Command == f.Command {
                        f.Run(b, &update.Message, update.Args)

                    }
                }
            }
        }
    }
}
