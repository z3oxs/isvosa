// A performatic Telegram API to build bots
package isvosa

import (
    "net/http"
    "fmt"
    "encoding/json"
    "io"
    "log"
    "strings"
)

func (b *Bot) Start() {
    var updates Updates
    var update Update
    
    for {
        r, err := http.Get(fmt.Sprintf("%s/bot%s/getUpdates?offset=-1&timeout=100", baseURL, b.Token))
        if err != nil {
            log.Fatal(err)

        } else if r.StatusCode != 200 {
            log.Fatal("Invalid token")

        }
        
        defer r.Body.Close()

        bytes, _ := io.ReadAll(r.Body)
        json.Unmarshal(bytes, &updates)
        update = updates.Update[len(updates.Update) - 1]
        
        if previousID == 0 {
            previousID = update.ID
        
            continue

        } else if previousID != update.ID {
            previousID = update.ID
            
            if update.Message.Text != "" && string(update.Message.Text[0]) == "/" && !update.Message.From.IsBot {
                rawText := strings.Split(update.Message.Text, " ")
                update.Command = rawText[0][1:]
                update.Args = rawText[1:]

                for _, c := range handler.Commands {
                    if update.Command == c.Command {
                        c.Run(b, &update.Message, update.Args)

                    }
                }
            }
        }
    }
}
