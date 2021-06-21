package app

import (
	"fmt"
	"math"
	"time"

	"github.com/leekchan/accounting"
)

func CalculateCapital() *CapitalResult {
	startTime := time.Date(2021, time.June, 1, 0, 0, 0, 0, time.Local)
	// capital := 1300000 // 03.2019
	capital := 1745000 // 05.2021
	salary := 284000   // 05.2021
	yearlyInvestmentProfitPercent := 10
	salaryBonusMonths := map[time.Month]bool{time.April: true, time.November: true}
	salaryBonusPercents := 10
	salaryGrowMonths := map[time.Month]bool{time.April: true, time.November: true}
	salaryGrowPercents := 5
	vacationMonths := map[time.Month]bool{time.December: true}
	vacationPrice := 300000
	goalAddition := 35000000
	goalMonthlyAddition := 100000

	monthlySpendingList := map[string]int{
		"квартира":    45000,
		"еда":         20000,
		"развлечения": 30000,
		"одежда":      5000,
		"путешествия": 30000,
	}

	var monthlySpending int
	for _, spending := range monthlySpendingList {
		monthlySpending += spending
	}

	goal := countGoal(goalAddition, goalMonthlyAddition, monthlySpending, yearlyInvestmentProfitPercent)

	calculator := accounting.Accounting{Symbol: "₽", Thousand: " ", Format: "%v %s"}

	capitalResult := NewCapitalResult(startTime, capital, salary, monthlySpending, goal, yearlyInvestmentProfitPercent)
	capitalResult.Append(
		[]string{
			"Дата",
			"Заработано",
			"Премия",
			"Инвестиции",
			"Заработано всего",
			"Капитал",
			"Отпуск",
		},
	)

	var overallProfit int
	var currentTime time.Time
	previousYear := startTime.Year()
	monthsToGoalCount := 0
	for ; capital < goal; monthsToGoalCount++ {
		currentTime = startTime.AddDate(0, monthsToGoalCount, 0)
		profit := salary - monthlySpending
		investmentProfit := int(float64(capital) * (float64(yearlyInvestmentProfitPercent) / 10 / 100))
		overallProfit += profit + investmentProfit

		if previousYear != currentTime.Year() {
			capitalResult.Append([]string{currentTime.Format("2006")})
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

		capitalResult.Append(row)
	}

	yearsToGoal, monthsToGoal := countTimeToGoal(monthsToGoalCount)

	capitalResult.GoalDate = startTime.AddDate(yearsToGoal, monthsToGoal, 0)
	capitalResult.ResultCapital = capital
	capitalResult.ResultSalary = salary

	return capitalResult
}

// count goal in terms of all spending = passive income
func countGoal(
	goalAddition int,
	goalMonthlyAddition int,
	monthlySpending int,
	yearlyInvestmentProfitPercent int,
) int {
	goal := float64(goalAddition) +
		(float64(monthlySpending)+float64(goalMonthlyAddition))/(float64(yearlyInvestmentProfitPercent)/10/100)

	return int(math.Round(goal))
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
