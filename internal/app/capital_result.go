package app

type CapitalResult struct {
	Table [][]string
}

func NewCapitalResult() *CapitalResult {
	return &CapitalResult{
		Table: make([][]string, 0),
	}
}

func (c *CapitalResult) Append(row []string) {
	c.Table = append(c.Table, row)
}
