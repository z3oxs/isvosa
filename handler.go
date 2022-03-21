package isvosa

func Add(command Command) {
    handler.Commands = append(handler.Commands, command)
}

func (b *Bot) Add(command string, run func(bot *Bot, msg *Message, args []string)) {
    handler.Commands = append(handler.Commands, Command { Command: command, Run: run })
}

func Commands() []Command {
    return handler.Commands
}
