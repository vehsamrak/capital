package renderer

import (
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/vehsamrak/capital/internal/app"
)

type Console struct {
}

func (c Console) Render(capitalResult *app.CapitalResult) {
	table := tablewriter.NewWriter(os.Stdout)

	capitalResultTable := capitalResult.Table

	for i, row := range capitalResultTable {
		if i == 0 {
			table.SetHeader(capitalResultTable[0])
			continue
		}

		table.Append(row)
	}

	table.Render()
}
