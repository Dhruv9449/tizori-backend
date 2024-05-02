package tizoriCli

import (
	"github.com/GDGVIT/Tizori-backend/cli/commands"
	"github.com/urfave/cli/v2"
)

func NewCliApp() *cli.App {
	// New cli app
	cliApp := cli.NewApp()

	// Set the name, usage and version of the app
	cliApp.Name = "Tizori"
	cliApp.Usage = "Tizori Backend API"
	cliApp.Version = "0.0.1"
	cliApp.Authors = []*cli.Author{
		{
			Name:  "Dhruv Shah",
			Email: "dhruvshahrds@gmail.com",
		}}
	cliApp.EnableBashCompletion = true
	cliApp.Description = "CLI for Tizori Backend API"

	// Set the commands
	commands.AddCommands(cliApp)

	return cliApp
}
