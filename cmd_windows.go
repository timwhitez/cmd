package cmd

import (
	"os/exec"
	"syscall"

	"github.com/shirou/gopsutil/process"
)

// Stop stops the command by sending its process group a SIGTERM signal.
// Stop is idempotent. An error should only be returned in the rare case that
// Stop is called immediately after the command ends but before Start can
// update its internal state.
func tP(pid int) error {
	p, _ := process.NewProcess(int32(pid)) // Specify process id of parent
	// handle error
	vp, _ := p.Children()
	var v *process.Process
	for _, v = range vp {
		_ = v.Kill() // Kill each child
		// handle error
	}
	return p.Kill()
}

func terminateProcess(pid int) error {
	//p := &os.Process{Pid: pid}
	//return p.Kill()
	return tP(pid)
}

func setProcessGroupID(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
