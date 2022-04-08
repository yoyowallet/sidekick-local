package main

import (
	"context"
	"os"
	"os/exec"
)

type Process struct {
	*exec.Cmd
	configSources []ConfigSource
}

func NewProcess(name string, arg ...string) *Process {
	p := new(Process)
	p.Cmd = exec.Command(name, arg...)

	// Defaults
	p.Stdin = os.Stdin
	p.Stdout = os.Stdout
	p.Stderr = os.Stderr

	return p
}

func (p *Process) AppendConfigSource(sources ...ConfigSource) {
	p.configSources = append(p.configSources, sources...)
}

func (p *Process) resetAndInstallEnv() error {
	// Reset
	p.Env = os.Environ()

	for _, configSource := range p.configSources {
		items, err := configSource.List(context.Background())
		if err != nil {
			return err
		}
		p.Env = append(p.Env, items...)
	}
	return nil
}

func (p *Process) Start() error {
	var err error

	err = p.resetAndInstallEnv()
	if err != nil {
		return err
	}

	return p.Cmd.Start()
}

func (p *Process) Stop() error {
	return p.Cmd.Process.Signal(os.Interrupt)

}

func (p *Process) Wait() error {
	return p.Cmd.Wait()
}
