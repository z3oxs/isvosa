package isvosa

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

func get(token, endpoint string, parsedStruct interface{}) []byte {
    data, _ := json.Marshal(parsedStruct)
    r, err := http.Post(fmt.Sprintf("%s/bot%s/%s", baseURL, token, endpoint), "application/json", bytes.NewBuffer(data))
    if err != nil {
        log.Fatal(err)

    } else if r.StatusCode != 200 {
        log.Fatal("Invalid token or endpoint")

    }

    defer r.Body.Close()
    
    bytes, _ := io.ReadAll(r.Body)
    return bytes
}

func (b *Bot) GetUserProfilePhotos(method GetUserProfilePhotos) UserProfilePhotos {
    r := get(b.Token, "getUserProfilePhotos", method)

    var photos UserProfilePhotos
    json.Unmarshal(r, &photos)

    return photos
}

func (b *Bot) GetChat(chatID int) GetChatResult {
    r := get(b.Token, "getChat", GetChat { chatID })

    var chat GetChatResult
    json.Unmarshal(r, &chat)

    return chat
}

func (b *Bot) GetChatAdministrators(chatID int) []ChatMemberAdministrator {
    r := get(b.Token, "getChatAdministrators", GetChat { chatID })

    var admins GetChatAdministratorsResult
    json.Unmarshal(r, &admins)

    return admins.Admins
}

func (b *Bot) GetChatMemberCount(chatID int) int {
    r := get(b.Token, "getChatMemberCount", GetChat { chatID })
    
    var count GetChatMemberCount
    json.Unmarshal(r, &count)

    return count.Count
}

func (b *Bot) GetMyCommands(scope string) []BotCommand {
    r := get(b.Token, "getMyCommands", BotCommandsScope { Type: scope })

    var botCommand GetMyCommandsResult
    json.Unmarshal(r, &botCommand)
    
    return botCommand.Commands
}

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
