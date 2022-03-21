package isvosa

// Append the commands parsed by parameters, to the array of Command of handler
func Add(command Command) {
    handler.Commands = append(handler.Commands, command)
}


// Do the same stuff above but directly using a bot instace and receiving the
// command treatment and a function that will handling the command
func (b *Bot) Add(command string, run func(bot *Bot, msg *Message, args []string)) {
    handler.Commands = append(handler.Commands, Command {
        Command: command,
        Run: run,
    })
}


// Will return the available and valid commands array
func Commands() []Command {
    return handler.Commands
}
