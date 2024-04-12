package main

import "github.com/urfave/cli/v2"

var metadataCmd = cli.Command{
	Name:        "metadata",
	Usage:       "Print metadata from parquet file",
	ArgsUsage:   "metadata [--format=table] <file>",
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
