package main

import (
	"fmt"
	"os"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
)

func main() {
	startTime := time.Date(2021, time.June, 1, 0, 0, 0, 0, time.Local)
	// capital := 1300000.0 // 03.2019
	capital := 1745000.0 // 05.2021
	salary := 284000.0   // 05.2021
	yearlyInvestmentProfitPercent := 10.0
	salaryBonusMonths := map[time.Month]bool{time.April: true, time.November: true}
	salaryBonusPercents := 10.0
	salaryGrowMonths := map[time.Month]bool{time.April: true, time.November: true}
	salaryGrowPercents := 5.0
	vacationMonths := map[time.Month]bool{time.December: true}
	vacationPrice := 300000.0
	goalAddition := 35000000.0
	goalMonthlyAddition := 100000.0

	monthlySpendingList := map[string]float64{
		"квартира":    45000,
		"еда":         20000,
		"развлечения": 30000,
		"одежда":      5000,
		"путешествия": 30000,
	}

	var monthlySpending float64
	for _, spending := range monthlySpendingList {
		monthlySpending += spending
	}

	goal := countGoal(goalAddition, goalMonthlyAddition, monthlySpending, yearlyInvestmentProfitPercent)

	// processing
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Дата", "Заработано", "Премия", "Инвестиции", "Заработано всего", "Капитал", "Отпуск"})

	calculator := accounting.Accounting{Symbol: "₽", Thousand: " ", Format: "%v %s"}

	fmt.Printf("Дата начала: %s\n", startTime.Format("2 January 2006"))
	fmt.Printf("Начальный капитал: %s\n", calculator.FormatMoney(capital))
	fmt.Printf("Начальная зарплата: %s\n", calculator.FormatMoney(salary))
	fmt.Printf(
		"Цель: %s (%s в месяц)\n",
		calculator.FormatMoney(goal),
		calculator.FormatMoney(goal*yearlyInvestmentProfitPercent/10/100),
	)
	fmt.Printf("Ежемесячные траты: %s\n", calculator.FormatMoney(monthlySpending))

	var overallProfit float64
	var currentTime time.Time
	previousYear := startTime.Year()
	monthsToGoalCount := 0
	for ; capital < goal; monthsToGoalCount++ {
		currentTime = startTime.AddDate(0, monthsToGoalCount, 0)
		profit := salary - monthlySpending
		investmentProfit := capital * (yearlyInvestmentProfitPercent / 10 / 100)
		overallProfit += profit + investmentProfit

		if previousYear != currentTime.Year() {
			table.Append([]string{currentTime.Format("2006")})
		}

		previousYear = currentTime.Year()

		capital += profit + investmentProfit

		var row []string
		row = append(
			row,
			currentTime.Format("January"),
			calculator.FormatMoney(profit),
			"",
			calculator.FormatMoney(investmentProfit),
			calculator.FormatMoney(overallProfit),
			calculator.FormatMoney(capital),
			"",
		)

		// vacation
		if _, ok := vacationMonths[currentTime.Month()]; ok {
			profit -= vacationPrice
			overallProfit -= vacationPrice
			capital -= vacationPrice
			row[1] = calculator.FormatMoney(profit)
			row[4] = calculator.FormatMoney(overallProfit)
			row[5] = calculator.FormatMoney(capital)
			if row[6] != "" {
				row[6] += " | "
			}
			row[6] = fmt.Sprintf("Отпуск: %s", calculator.FormatMoney(vacationPrice))
		}

		// salary growth
		if _, ok := salaryGrowMonths[currentTime.Month()]; ok {
			salaryGrow := salary * salaryGrowPercents / 100
			salary += salaryGrow
			if row[6] != "" {
				row[6] += " | "
			}
			row[6] = fmt.Sprintf(
				"Повышение зарплаты: %s (+%s)",
				calculator.FormatMoney(salary),
				calculator.FormatMoney(salaryGrow),
			)
		}

		// salary bonus
		if _, ok := salaryBonusMonths[currentTime.Month()]; ok {
			bonus := salary * salaryBonusPercents / 100 * 6
			profit += bonus
			overallProfit += bonus
			row[1] = calculator.FormatMoney(profit)
			row[2] += calculator.FormatMoney(bonus)
			row[4] = calculator.FormatMoney(overallProfit)
		}

		table.Append(row)
	}

	yearsToGoal, monthsToGoal := countTimeToGoal(monthsToGoalCount)
	fmt.Printf("До цели: %d лет %d месяцев\n", yearsToGoal, monthsToGoal)

	table.Render()
}

// count goal in terms of all spending = passive income
func countGoal(
	goalAddition float64,
	goalMonthlyAddition float64,
	monthlySpending float64,
	yearlyInvestmentProfitPercent float64,
) float64 {
	return goalAddition +
		((monthlySpending + goalMonthlyAddition) / (yearlyInvestmentProfitPercent / 10 / 100))
}

func countTimeToGoal(monthsToGoal int) (int, int) {
	if monthsToGoal == 0 {
		return 0, 0
	}

	if monthsToGoal <= 12 {
		return 0, monthsToGoal
	}

	return monthsToGoal / 12, monthsToGoal % 12
}
