package isvosa

import (
    "fmt"
    "strings"
)

func (b *Bot) Send(method interface{}) {
    request(method, b.Token, strings.Split(fmt.Sprintf("%T", method), ".")[1])
}
