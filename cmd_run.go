package main

import "github.com/urfave/cli"

var commandRun = cli.Command{
	Name: "run",
	Action: func(c *cli.Context) error {
		var err error

		if c.NArg() == 0 {
			return cli.NewExitError("missing command to run", 1)
		}

		proc := NewProcess(c.Args().First(), c.Args().Tail()...)
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
