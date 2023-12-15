package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/snowboardit/reserved/pkg/reserved"
	"github.com/urfave/cli/v2"
)

const (
	name    = "reserved"
	version = "v0.0.1"
)

var (
	App *cli.App
	r   = reserved.Reserved{}
)

func Init() {
	App = &cli.App{
		Name:     name,
		Version:  version,
		Compiled: time.Now(),
		// Authors: []*cli.Author{
		// 	{
		// 		Name:  "Max Lareau",
		// 		Email: "info@maxlareau.com",
		// 	},
		// },
		HelpName:  "Reserved",
		Usage:     "a cli tool for determining if word(s) are reserved in programming/database languages",
		UsageText: "$ reserved [options] [words...]",
		ArgsUsage: "[words, names, and such]",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "programming", Aliases: []string{"p"}},
			&cli.BoolFlag{Name: "database", Aliases: []string{"d"}},
			&cli.BoolFlag{Name: "stack", Aliases: []string{"s"}},
		},
		EnableBashCompletion: true,
		Suggest:              true,
		HideHelp:             false,
		// HideVersion:          false,
		Action: func(cCtx *cli.Context) error {
			// get arg(s)
			args := cCtx.Args().Slice()

			// show help if no args
			if len(args) < 1 {
				cli.ShowAppHelp(cCtx)
				return nil
			}

			// stack flag
			if cCtx.Bool("stack") {
				fmt.Printf("Args: %s", strings.Join(args, ", "))
				fmt.Printf("Checking if words are reserved in stack %s...", "stack")
				fmt.Println()
				return nil
			}

			// programming flag
			if cCtx.Bool("programming") {
				fmt.Printf("Args: %s", strings.Join(args, ", "))
				fmt.Println("Checking if words are reserved in programming languages...")
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
			if len(words.Reserved) > 0 {
				fmt.Println()
				fmt.Println("❗️ Words are reserved")
				fmt.Printf("%s", words.String())
				fmt.Println()
			} else {
				fmt.Println("✅ None of the words are reserved")
			}

			return nil
		},
	}
}

func Start() {
	// Cli setup
	Init()

	// Run the cli
	App.Run(os.Args)
}
