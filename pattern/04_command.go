package pattern

type Command interface {
	Execute() string
}

type ToggleOnCommand struct {
	receiver *Receiver
}

func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

type ToggleOffCommand struct {
	receiver *Receiver
}

func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

type Receiver struct {
}

func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}
