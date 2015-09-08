package executor

import (
	"os/exec"
)

type Executor struct {
	memory map[int]exec.Cmd
}

func NewExecutor() *Executor {
	return &Executor{memory: make(map[int]exec.Cmd)}
}

func (e *Executor) Start(path string) (pid int, err error) {
	command := exec.Cmd{Path: path}
	if err = command.Start(); err != nil {
		return
	}

	pid = command.Process.Pid
	e.memory[pid] = command

	return
}
