package command

import (
	log "github.com/Sirupsen/logrus"
	"github.com/TuringCodeOne/dockerdemo/container"
	"github.com/urfave/cli"
)



var initCommand = cli.Command {
	Name: "init",
	Usage: `Init container process run user's process in container. Don't call it outside`,
	Action: func(context *cli.Context) error {
		log.Info("Container init...")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
