package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/snowboardit/reserved/pkg/cli/commands"
	"github.com/snowboardit/reserved/pkg/reserved"
	"github.com/urfave/cli/v2"
)

const (
	name                = "reserved"
	version             = "v0.0.1"
	wordsAreReserved    = "❌ Some of the words are reserved"
	wordsAreNotReserved = "✅ None of the words are reserved"
)

var (
	App *cli.App
)

func init() {
	// setup app
	App = &cli.App{
		Name:      name,
		Version:   version,
		Compiled:  time.Now(),
		HelpName:  "Reserved",
		Usage:     "a cli to determine if word(s) are reserved in programming/database languages",
		UsageText: "$ reserved [options] [words...]",
		ArgsUsage: "[words, names, and such]",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "programming", Aliases: []string{"p"}},
			&cli.BoolFlag{Name: "database", Aliases: []string{"d"}},
		},
		EnableBashCompletion: true,
		Suggest:              true,
		HideHelp:             false,
		Action: func(cCtx *cli.Context) error {
			// get arg(s)
			args := cCtx.Args().Slice()

			// show help if no args
			if len(args) < 1 {
				cli.ShowAppHelp(cCtx)
				return nil
			}

			var r = reserved.New()

			// stack flag
			if cCtx.Bool("stack") {
				fmt.Printf("Args: %s", strings.Join(args, ", "))
				fmt.Printf("Checking if words are reserved in stack %s...", "stack")
				fmt.Println()
				return nil
			}

			// programming flag
			if cCtx.Bool("programming") {
				words := r.CheckProgramming(args...)
				if len(words) > 0 {
					fmt.Println(wordsAreReserved)
					fmt.Println()
					fmt.Printf("%s", words.String())
					fmt.Println()
					fmt.Println()
				} else {
					fmt.Println(wordsAreNotReserved)
				}
				return nil
			}

			// database flag
			if cCtx.Bool("database") {
				fmt.Printf("Args: %s", strings.Join(args, ", "))
				fmt.Println("Checking if words are reserved in database languages...")
				return nil
			}

			// default: check if args are reserved across all languages
			words := r.Check(args...)
			if len(words) > 0 {
				fmt.Println(wordsAreReserved)
				fmt.Println()
				fmt.Printf("%s", words.String())
				fmt.Println()
				fmt.Println()
			} else {
				fmt.Println(wordsAreNotReserved)
			}

			return nil
		},
	}

	// setup commands
	commands.Init(App)
}

func Start() {
	// Run
	App.Run(os.Args)
}
