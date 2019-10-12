package main

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/TuringCodeOne/dockerdemo/command"
	"github.com/urfave/cli"
	"os"
)

const usage = `dockerdemo is a simple container runtime implementation,
			   The purpose of this project is to learn how docker works and how to write a docker by ourselves,
               Enjoy it, just for fun!`

func main() {
	app := cli.NewApp()
	app.Name = "dockerdemo"
	app.Usage = usage

	app.Commands = []cli.Command {
		InitCommand,
		RunCommand,
	}

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdin)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}


