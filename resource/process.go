package resource

import "github.com/aelsabbahy/goss/system"

type Process struct {
	Executable string `json:"executable"`
	Running    bool   `json:"running"`
}

func (p *Process) Validate(sys *system.System) []TestResult {
	sysProcess := sys.NewProcess(p.Executable, sys)

	var results []TestResult

	results = append(results, ValidateValue(p.Executable, "running", p.Running, sysProcess.Running))

	return results
}

func NewProcess(sysProcess system.Process) *Process {
	executable := sysProcess.Executable()
	running, _ := sysProcess.Running()
	return &Process{
		Executable: executable,
		Running:    running.(bool),
	}
}