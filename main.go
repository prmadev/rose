package main

import (
	"fmt"
	"log"
	"os"
	"rose/match"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v2"
)

var prFlags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:    "colorname",
		Aliases: []string{"c"},
		Value:   "",
		Usage:   "Color name.",
	},
	&cli.StringFlag{
		Name:    "swatch",
		Aliases: []string{"s"},
		Value:   "rp",
		Usage:   "Swatch to use.",
	},
	&cli.StringFlag{
		Name:    "palette",
		Aliases: []string{"p"},
		Value:   "rosePine",
		Usage:   "Palette to use.(not developed for now)",
	},
}

func main() {

	app := &cli.App{
		Name:  "rose",
		Usage: "Your friendly rosepine helper.",
		Commands: []*cli.Command{
			{
				Name:    "pr",
				Aliases: []string{"p"},
				Usage:   "Print out the color code of the given string",
				Flags:   prFlags,
				Action:  printer,
			},
			{
				Name:    "copy",
				Aliases: []string{"c"},
				Usage:   "Store code inside clipboard",
				Flags:   prFlags,
				Action:  cp,
			},
			{
				Name:    "print-all",
				Aliases: []string{"a"},
				Usage:   "print all the codes of a colorscheme",
				Flags:   prFlags,
				Action:  printAll,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
func printer(c *cli.Context) error {
	col := match.Get(c.String("colorname"), c.String("swatch"))
	block := fmt.Sprint(lipgloss.NewStyle().Background(lipgloss.Color(col)).PaddingRight(2).Render(""))
	code := fmt.Sprint(lipgloss.NewStyle().Bold(true).Italic(true).Render(block + " " + col))
	fmt.Println(code)
	return nil
}
func cp(c *cli.Context) error {
	col := match.Get(c.String("colorname"), c.String("swatch"))
	err := clipboard.WriteAll(col)
	if err != nil {
		return err
	}
	return nil
}
func printAll(c *cli.Context) error {
	swatch := match.List(c.String("swatch"))
	for n, ing := range swatch {
		col := ing.Hex
		block := fmt.Sprint(lipgloss.NewStyle().Background(lipgloss.Color(col)).PaddingRight(2).Render(""))
		code := fmt.Sprint(lipgloss.NewStyle().Bold(true).Italic(true).Render(block + " " + col + " " + n))
		fmt.Println(code)
	}
	return nil
}
