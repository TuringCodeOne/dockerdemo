package command

import (
	"github.com/Sirupsen/logrus"
	"github.com/TuringCodeOne/dockerdemo/cgroup"
	"github.com/TuringCodeOne/dockerdemo/cgroup/subsystem"
	"github.com/TuringCodeOne/dockerdemo/container"
	"os"
	"strings"
)

func Run(tty bool, comArray []string, res *subsystem.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		logrus.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}

	cgroupMngr := cgroup.NewCgroupManager("dockerdemo-cgroup1")
	defer cgroupMngr.Destroy()
	_ = cgroupMngr.Set(res)
	_ = cgroupMngr.Apply(parent.Process.Pid)
	sendInitCommand(comArray, writePipe)
	_ = parent.Wait()
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	logrus.Infof("command all is %s", command)
	_, _ = writePipe.WriteString(command)
	writePipe.Close()
}
