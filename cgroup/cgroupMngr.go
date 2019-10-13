package cgroup

import (
	"github.com/Sirupsen/logrus"
	"github.com/TuringCodeOne/dockerdemo/cgroup/subsystem"
)

type Manager struct {
	Path string
	Resource *subsystem.ResourceConfig
}

func NewCgroupManager(path string) *Manager {
	return &Manager{
		Path:     path,
		Resource: nil,
	}
}

func (c *Manager) Apply(pid int) error {
	for _, subSysIns := range subsystem.SubsysIns {
		_ = subSysIns.Apply(c.Path, pid)
	}
	return nil
}

func (c *Manager) Set(res *subsystem.ResourceConfig) error {
	for _, subSysIns := range subsystem.SubsysIns {
		_ = subSysIns.Set(c.Path, res)
	}
	return nil
}

func (c *Manager) Destroy() error {
	for _, subsSysIns := range subsystem.SubsysIns {
		if err := subsSysIns.Remove(c.Path); err != nil {
			logrus.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}

