package easycmd

type Command struct {
	Name        string
	Aliases     []string
	Description string
	Usage       string
	Run         func(args []string)
}

func NewCommand(name string) Command {
	return Command{Name: name}
}
