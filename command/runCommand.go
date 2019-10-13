package command

import (
	"fmt"
	"github.com/TuringCodeOne/dockerdemo/cgroup/subsystem"
	"github.com/urfave/cli"
)

var RunCommand = cli.Command {
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit, dockerdemo run -it [command]`,
	Flags: []cli.Flag {
		cli.BoolFlag {
			Name:        "it",
			Usage:       "enable tty",
		},
		cli.StringFlag {
			Name:        "m",
			Usage:       "memory limit",
		},
		cli.StringFlag {
			Name:        "cpushare",
			Usage:       "cpushare limit",
		},
		cli.StringFlag{
			Name:        "cpuset",
			Usage:       "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}
		tty := context.Bool("it")
		resConf := &subsystem.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuShare:    context.String("cpuset"),
			CpuSet:      context.String("cpushare"),
		}

		Run(tty, cmdArray, resConf)
		return nil
	},
}
