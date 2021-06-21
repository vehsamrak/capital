package app

import "time"

type CapitalResult struct {
	InitialDate     time.Time  `json:"startDate"`
	InitialCapital  int        `json:"initialCapital"`
	ResultCapital   int        `json:"resultCapital"`
	InitialSalary   int        `json:"initialSalary"`
	ResultSalary    int        `json:"resultSalary"`
	Goal            int        `json:"goal"`
	MonthlyGoal     int        `json:"monthlyGoal"`
	GoalDate        time.Time  `json:"goalDate"`
	MonthlySpending int        `json:"monthlySpending"`
	Table           [][]string `json:"table"`
}

func NewCapitalResult(
	startTime time.Time,
	capital int,
	salary int,
	monthlySpending int,
	goal int,
	yearlyInvestmentProfitPercent int,
) *CapitalResult {
	return &CapitalResult{
		InitialDate:     startTime,
		InitialCapital:  capital,
		ResultCapital:   capital,
		InitialSalary:   salary,
		ResultSalary:    salary,
		Goal:            goal,
		MonthlyGoal:     goal * yearlyInvestmentProfitPercent / 10 / 100,
		MonthlySpending: monthlySpending,
		Table:           make([][]string, 0),
	}
}

func (c *CapitalResult) SetGoalDate(goalDate time.Time) {
	c.GoalDate = goalDate
}

func (c *CapitalResult) Append(row []string) {
	c.Table = append(c.Table, row)
}
