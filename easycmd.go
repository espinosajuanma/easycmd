package easycmd

type App struct {
	Commands    []Command
	Description string
	Author      string
}

func (r *App) AddCommand(cmd Command) []Command {
	r.Commands = append(r.Commands, cmd)
	return r.Commands
}

func (r App) Run(command string, args []string) {
	for _, cmd := range r.Commands {
		if cmd.Name == command || cmd.IsAlias(command) {
			cmd.Run(args)
		}
	}
}

func NewApp(description string, author string) App {
	app := App{
		[]Command{},
		description,
		author,
	}
	return app
}
