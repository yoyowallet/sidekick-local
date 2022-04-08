package main

import (
	"os"

	"github.com/google/shlex"
	yaml "gopkg.in/yaml.v2"
)

type Procfile struct {
	Commands map[string][]string
}

func readProcfile(name string) (*Procfile, error) {
	var err error

	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	dec.SetStrict(true)

	var entries map[string]string
	err = dec.Decode(&entries)
	if err != nil {
		return nil, err
	}

	pf := new(Procfile)
	pf.Commands = make(map[string][]string, len(entries))

	for processType, command := range entries {
		parsedCommand, err := shlex.Split(command)
		if err != nil {
			return nil, err
		}
		pf.Commands[processType] = parsedCommand
	}

	return pf, nil
}
