package cmd

import (
	"github.com/urfave/cli"
	"os"
)

type Param struct {
	ConfPath         string
	FetchInterval    string
	RegisterInterval string
}

var CLIParam = &Param{}

func ReadParams() error {
	app := cli.NewApp()
	app.Name = "go-chassis benchmark tool"
	app.Description = "example: ./main -c /etc/registrator/reg.yaml"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "c",
			Value:       "./conf/reg.yaml",
			Usage:       "config path",
			Destination: &CLIParam.ConfPath,
		},
	}
	app.Action = func(c *cli.Context) error {
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return nil
}
