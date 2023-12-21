package commands

import (
	"github.com/snowboardit/reserved/pkg/reserved"

	"fmt"

	"github.com/urfave/cli/v2"
)

var commands []*cli.Command = []*cli.Command{
	{
		Name:        "languages",
		Usage:       "list all loaded languages",
		Description: "list all loaded languages",
		Action: func(cCtx *cli.Context) error {
			r := reserved.New()
			languages := r.Languages()

			if len(languages) < 1 {
				return fmt.Errorf("No languages found")
			}

			for _, w := range languages {
				fmt.Printf("%s\n", w)
			}

			return nil
		},
	},
}

func Init(a *cli.App) {
	a.Commands = append(a.Commands, commands...)
}
