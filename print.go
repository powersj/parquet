package main

import "github.com/urfave/cli/v2"

var printCmd = cli.Command{
	Name:        "print",
	Usage:       "Print parquet file",
	ArgsUsage:   "print [--format=table] <file>",
	Description: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "output format, either 'table' (default) or 'json'",
		},
	},
	Action: func(_ *cli.Context) error {
		return nil
	},
}
