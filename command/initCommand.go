package command

import (
	log "github.com/Sirupsen/logrus"
	"github.com/TuringCodeOne/dockerdemo/container"
	"github.com/urfave/cli"
)



var InitCommand = cli.Command {
	Name: "init",
	Usage: `Init container process run user's process in container. Don't call it outside`,
	Action: func(context *cli.Context) error {
		log.Info("Container init...")
		err := container.RunContainerInitProcess()
		return err
	},
}
