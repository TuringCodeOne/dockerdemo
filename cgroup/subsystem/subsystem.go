package subsystem

type ResourceConfig struct {
	MemoryLimit string
	CpuShare	string
	CpuSet		string
}

type SubSystem interface {
	Name() string
	Set(Path string, res *ResourceConfig) error
	Apply(path string, pid int) error
	Remove(path string) error
}

var (
	SubsysIns = []SubSystem {
		&CpusetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSystem{},
	}
)
