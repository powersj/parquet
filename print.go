package main

import (
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v16/parquet/file"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/powersj/parquet/internal"
	"github.com/urfave/cli/v2"
)

var printCmd = cli.Command{
	Name:        "print",
	Usage:       "Print parquet file",
	ArgsUsage:   "print [--format=table] <file>",
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
		data := make([]map[string]any, 0, metadata.NumRows)
		columns := make([]string, metadata.Schema.NumColumns())
		for i := 0; i < reader.NumRowGroups(); i++ {
			rowGroup := reader.RowGroup(i)
			scanners := make([]*internal.ColumnParser, metadata.Schema.NumColumns())
			for colIndex := range metadata.Schema.NumColumns() {
				col, err := rowGroup.Column(colIndex)
				if err != nil {
					return fmt.Errorf("unable to fetch column %q: %w", colIndex, err)
				}

				scanners[colIndex] = internal.NewColumnParser(col)
				columns[colIndex] = col.Descriptor().Name()
			}

			rowIndex := 0
			dataGroup := make([]map[string]any, rowGroup.NumRows())
			for _, s := range scanners {
				for s.HasNext() {
					if rowIndex%int(rowGroup.NumRows()) == 0 {
						rowIndex = 0
					}

					val, ok := s.Next()
					if !ok || val == nil {
						rowIndex++
						continue
					}

					if dataGroup[rowIndex] == nil {
						dataGroup[rowIndex] = make(map[string]any, metadata.Schema.NumColumns())
					}

					dataGroup[rowIndex][s.Name] = val
					rowIndex++
				}
			}

			data = append(data, dataGroup...)
		}

		switch ctx.String("format") {
		case "json":
			data, _ := json.MarshalIndent(data, "", "  ")
			if string(data) == "null" {
				fmt.Println("[]")
			} else {
				fmt.Println(string(data))
			}
		case "table":
			dataTable := table.NewWriter()

			var headerRow table.Row
			for _, c := range columns {
				headerRow = append(headerRow, c)
			}
			dataTable.AppendHeader(headerRow)

			for _, row := range data {
				var dataRow table.Row
				for _, c := range columns {
					dataRow = append(dataRow, row[c])
				}
				dataTable.AppendRow(dataRow)
			}

			dataTable.AppendFooter(table.Row{"total", fmt.Sprintf("%d rows", len(data))})

			fmt.Println(dataTable.Render())
		}

		return nil
	},
}
