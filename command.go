package easycmd

type Command struct {
	Name    string
	Aliases []string
	Usage   string
	Run     func(args []string)
}

func NewCommand(name string, aliases []string, usage string, Run func(args []string)) Command {
	return Command{
		name, aliases, usage, Run,
	}
}

func (cmd Command) IsAlias(alias string) bool {
	for _, c := range cmd.Aliases {
		if c == alias {
			return true
		}
	}
	return false
}
