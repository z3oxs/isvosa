package isvosa

func (b *Bot) Ban(chatID, userID int) {
    b.Send(BanChatMember { ChatID: chatID, UserID: userID })
}

func (b *Bot) Unban(chatID, userID int) {
    b.Send(UnbanChatMember { ChatID: chatID, UserID: userID })
}
