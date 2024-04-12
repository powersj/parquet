package main

import "github.com/urfave/cli/v2"

var schemaCmd = cli.Command{
	Name:        "schema",
	Usage:       "Print schema of parquet file",
	ArgsUsage:   "schema [--format=table] <file>",
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
