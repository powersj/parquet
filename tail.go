package main

import "github.com/urfave/cli/v2"

var tailCmd = cli.Command{
	Name:        "tail",
	Usage:       "Print last N rows of parquet file",
	ArgsUsage:   "metadata [--format=table] [--lines=10] <file>",
	Description: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "output format, either 'table' (default) or 'json'",
		},
		&cli.IntFlag{
			Name:  "lines",
			Usage: "number of lines to print (default: 10)",
		},
	},
	Action: func(_ *cli.Context) error {
		return nil
	},
}
