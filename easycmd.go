package easycmd

type App struct {
	Commands    map[string]*Command
	Aliases     map[string]*Command
	Description string
	Author      string
	Version     string
}

func (app *App) AddCommand(cmd Command) {
	app.Commands[cmd.Name] = &cmd
	for _, alias := range cmd.Aliases {
		app.Aliases[alias] = &cmd
	}
}

func (app App) Run(command string, args []string) {
	if command == "" {
		app.Commands["help"].Run(args)
		return
	}

	cmd, ok := app.Commands[command]
	if ok {
		cmd.Run(args)
		return
	}

	alias, ok := app.Aliases[command]
	if ok {
		alias.Run(args)
		return
	}

	app.Commands["help"].Run(args)
}

func NewApp() *App {
	app := new(App)
	app.Commands = make(map[string]*Command)
	app.Aliases = make(map[string]*Command)

	return app
}
