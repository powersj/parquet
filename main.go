package main

import (
	"fmt"
	"os"

	"github.com/powersj/parquet/internal"
	"github.com/urfave/cli/v2"
)

const (
	cliDescription = `parquet is a CLI to interact with parquet files`
)

func main() {
	app := &cli.App{
		Name:        "parquet",
		Usage:       "CLI to interact with parquet files",
		Description: cliDescription,
		Suggest:     true,
		Commands: []*cli.Command{
			&metadataCmd,
			&printCmd,
			&queryCmd,
			&schemaCmd,
			{
				Name:        "version",
				Usage:       "Print version, build, and platform info",
				Description: "Print version, build, and platform info",
				Action: func(*cli.Context) error {
					fmt.Println(internal.AppVersion())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
