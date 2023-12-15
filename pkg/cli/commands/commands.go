package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var commands []*cli.Command = []*cli.Command{
	{
		Name: "reserved",
		// Aliases:     []string{"rd"},
		Category:    "base",
		Usage:       "$ reserved [options] [words...]",
		UsageText:   "doo - does the dooing",
		Description: "no really, there is a lot of dooing to be done",
		ArgsUsage:   "[words...]",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "forever", Aliases: []string{"forevvarr"}},
		},
		Action: func(cCtx *cli.Context) error {
			cCtx.Command.FullName()
			cCtx.Command.HasName("wop")
			cCtx.Command.Names()
			cCtx.Command.VisibleFlags()
			fmt.Fprintf(cCtx.App.Writer, "dodododododoodododddooooododododooo\n")
			if cCtx.Bool("forever") {
				cCtx.Command.Run(cCtx)
			}
			return nil
		},
	},
}

func Init(a *cli.App) {
	a.Commands = append(a.Commands, commands...)
}
