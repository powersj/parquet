package main

import (
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v16/parquet/file"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

type fileInfo struct {
	Version    string       `json:"version"`
	CreatedBy  string       `json:"created_by"`
	NumRows    int64        `json:"num_rows"`
	NumGroups  int          `json:"num_groups"`
	NumColumns int          `json:"num_columns"`
	Columns    []columnInfo `json:"columns"`
}

type columnInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

var infoCmd = cli.Command{
	Name:        "info",
	Usage:       "Print metadata and schema information about parquet file",
	ArgsUsage:   "info [--format=table] <file>",
	Description: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "output format, either 'table' (default) or 'json'",
			Value: "table",
		},
	},
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() != 1 {
			return fmt.Errorf("expected exactly one argument")
		}

		switch ctx.String("format") {
		case "json", "table":
		default:
			return fmt.Errorf("invalid format option: %s", ctx.String("format"))
		}

		filename := ctx.Args().First()
		reader, err := file.OpenParquetFile(filename, true)
		if err != nil {
			return err
		}
		defer reader.Close()

		metadata := reader.MetaData()
		fInfo := fileInfo{
			Version:    metadata.Version().String(),
			CreatedBy:  *metadata.CreatedBy,
			NumRows:    metadata.NumRows,
			NumGroups:  len(metadata.RowGroups),
			NumColumns: metadata.Schema.NumColumns(),
			Columns:    make([]columnInfo, metadata.Schema.NumColumns()),
		}
		for i := range metadata.Schema.NumColumns() {
			col := metadata.Schema.Column(i)
			fInfo.Columns[i] = columnInfo{
				Name: col.Name(),
				Type: col.PhysicalType().String(),
			}
		}

		switch ctx.String("format") {
		case "json":
			data, _ := json.MarshalIndent(fInfo, "", "  ")
			if string(data) == "null" {
				fmt.Println("[]")
			} else {
				fmt.Println(string(data))
			}
		case "table":
			fileTable := table.NewWriter()
			fileTable.AppendRow(table.Row{"filename", filename})
			fileTable.AppendRow(table.Row{"version", fInfo.Version})
			fileTable.AppendRow(table.Row{"created by", fInfo.CreatedBy})
			fileTable.AppendRow(table.Row{"number of rows", fInfo.NumRows})
			fileTable.AppendRow(table.Row{"number of groups", fInfo.NumGroups})
			fileTable.AppendRow(table.Row{"number of columns", fInfo.NumColumns})

			columns := table.NewWriter()
			columns.AppendHeader(table.Row{"column name", "type"})
			for _, column := range fInfo.Columns {
				columns.AppendRow(table.Row{
					column.Name,
					column.Type,
				})
			}
			fileTable.AppendRow(table.Row{"columns", columns.Render()})

			fmt.Println(fileTable.Render())
		}

		return nil
	},
}
