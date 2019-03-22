package main

import (
	"fmt"
	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	// configuration
	currentCapital := 1300000.0
	// currentCapital := 800000.0
	months := 18
	salary := 145000.0
	monthlyInvestmentProfitPercent := 3.0
	monthsToPremium := 6
	premiumSalaryPercent := 100.0
	monthsToSalaryGrow := 6
	salaryGrowInPercents := 4.0
	monthlySpendings := map[string]float64{
		"квартира":    45000,
		"еда":         11000,
		"развлечения": 20000,
	}

	// processing
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Месяц", "Заработано", "Премия", "Инвестиции", "Заработано всего", "Капитал"})

	calculator := accounting.Accounting{Symbol: "руб.", Thousand: " ", Format: "%v %s"}

	var monthlySpending float64
	for _, spending := range monthlySpendings {
		monthlySpending += spending
	}

	fmt.Printf("Начальный капитал: %s\n", calculator.FormatMoney(currentCapital))
	fmt.Printf("Начальная зарплата: %s\n", calculator.FormatMoney(salary))
	fmt.Printf("Ежемесячные траты: %s\n", calculator.FormatMoney(monthlySpending))

	var earnedCapital float64
	for i := 0; months > i; i++ {
		var row []string
		var earned float64
		earned += salary - monthlySpending

		investmentProfit := (earnedCapital + currentCapital) * (monthlyInvestmentProfitPercent / 100)

		row = append(
			row,
			fmt.Sprintf("%d", i+1),
			calculator.FormatMoney(earned),
			"",
			calculator.FormatMoney(investmentProfit),
			calculator.FormatMoney(earned+investmentProfit+earnedCapital),
			calculator.FormatMoney(earnedCapital+investmentProfit+currentCapital),
		)

		if (i+1)%(monthsToPremium) == 0 {
			var premium float64
			premium = salary * (premiumSalaryPercent / 100)
			earned += premium

			row[2] = calculator.FormatMoney(premium)
		}
		if (i+1)%(monthsToSalaryGrow) == 0 {
			salary *= (salaryGrowInPercents / 100) + 1
		}

		earnedCapital += earned + investmentProfit

		table.Append(row)
	}

	table.Render()
}
