package main

import (
	"github.com/urfave/cli"
)

var commandStart = cli.Command{
	Name: "start",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "Procfile",
		},
	},
	Action: func(c *cli.Context) error {
		var err error

		if c.NArg() != 1 {
			return cli.NewExitError("missing process type", 2)
		}
		processType := c.Args().Get(0)

		filename := c.String("file")
		pf, err := readProcfile(filename)
		if err != nil {
			return err
		}

		command, ok := pf.Commands[processType]
		if !ok {
			return cli.NewExitError("couldn't find that process type", 3)
		}

		proc := NewProcess(command[0], command[1:]...)
		proc.AppendConfigSource(configSourcesFromContext(c)...)

		err = proc.Start()
		if err != nil {
			return err
		}

		err = proc.Wait()
		if err != nil {
			return err
		}

		return nil
	},
}
