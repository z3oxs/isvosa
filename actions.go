package isvosa

func (b *Bot) StopPoll(chatID, messageID int) {
    request(StopPoll{ ChatID: chatID, MessageID: messageID }, b.Token, "stopPoll")
}

func (b *Bot) Delete(chatID, messageID int) {
    request(DeleteMessage{ ChatID: chatID, MessageID: messageID }, b.Token, "deleteMessage")
}
