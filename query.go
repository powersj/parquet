package main

import "github.com/urfave/cli/v2"

var queryCmd = cli.Command{
	Name:        "query",
	Usage:       "Query parquet file",
	ArgsUsage:   "query [--format=table] <file> <query>",
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
