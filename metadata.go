package main

import (
	"fmt"

	"github.com/apache/arrow/go/v16/parquet/file"
	"github.com/urfave/cli/v2"
)

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
	Action: func(ctx *cli.Context) error {
		filename := ctx.Args().First()
		reader, err := file.OpenParquetFile(filename, true)
		if err != nil {
			return err
		}
		defer reader.Close()

		// Print metadata
		metadata := reader.MetaData()

		fmt.Println(metadata.Version())
		fmt.Println(metadata.NumRows)
		fmt.Println(metadata.Schema)
		fmt.Println(metadata.EncryptionAlgorithm())

		fmt.Println("Schema")
		for i := range metadata.Schema.NumColumns() {
			f := metadata.Schema.Column(i)
			fmt.Println(f.Name())
			fmt.Println(f.PhysicalType())
			fmt.Println()
		}

		fmt.Println("rows:    ", reader.NumRows())
		fmt.Println("groups:  ", reader.NumRowGroups())
		fmt.Println("columns: ", reader.RowGroup(0).NumColumns())

		return nil
	},
}
