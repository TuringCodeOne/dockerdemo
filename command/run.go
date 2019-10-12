package command

import (
	"dockerdemo/container"
	"github.com/Sirupsen/logrus"
	"os"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}
	_ = parent.Wait()
	os.Exit(-1)
}