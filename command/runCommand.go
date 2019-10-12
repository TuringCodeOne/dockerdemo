package command

import (
	"fmt"
	"github.com/urfave/cli"
)

var runCommand = cli.Command {
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit,
			dockerdemo run -it [command]`,
	Flags: []cli.Flag {
		cli.BoolFlag {
			Name:        "it",
			Usage:       "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("it")
		Run(tty, cmd)
		return nil
	},
}
