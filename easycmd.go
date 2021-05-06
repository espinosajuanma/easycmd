package easycmd

type App struct {
	Commands []Command
}

func (r *App) AddCommand(cmd Command) []Command {
	r.Commands = append(r.Commands, cmd)
	return r.Commands
}

func (r App) Run(command string, args []string) {
	for _, cmd := range r.Commands {
		if cmd.Name == command || IsAlias(command, cmd.Aliases) {
			cmd.Run(args)
		}
	}
}

func NewApp() App {
	app := App{
		[]Command{},
	}
	return app
}

func IsAlias(alias string, listAlias []string) bool {
	for _, b := range listAlias {
		if b == alias {
			return true
		}
	}
	return false
}

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
