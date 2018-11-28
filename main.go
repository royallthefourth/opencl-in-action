package main

import (
	"github.com/urfave/cli"
	"log"
	"opencl-in-action/ch2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "opencl-in-action"
	app.Description = "Example programs from OpenCL In Action adapted to Go"

	app.Commands = []cli.Command{
		{
			Name:  "ch2",
			Usage: "Examples from chapter 2",
			Subcommands: []cli.Command{
				{Name: "list",
					Usage:  "List OpenCL extensions",
					Action: ch2.GetExtensions,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
