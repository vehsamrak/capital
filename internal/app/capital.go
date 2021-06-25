package app

import (
	"fmt"
	"math"
	"time"

	"github.com/leekchan/accounting"

	"github.com/vehsamrak/capital/internal/app/config"
)

func CalculateCapital(config *config.Config) *CapitalResult {
	startTime := config.StartTime
	capital := config.InitialCapital
	salary := config.InitialSalary
	yearlyInvestmentProfitPercent := config.YearlyInvestmentProfitPercent
	salaryBonusMonths := config.SalaryBonusMonths
	salaryBonusPercent := config.SalaryBonusPercent
	salaryGrowMonths := config.SalaryGrowthMonths
	salaryGrowPercent := config.SalaryGrowthPercent
	vacationMonths := config.VacationMonths
	vacationSpendingPercent := config.VacationSpendingPercent
	goalAddition := config.GoalAddition
	goalMonthlyAddition := config.GoalMonthlyAddition
	monthlySpendingList := config.MonthlySpending

	var monthlySpending int
	for _, spending := range monthlySpendingList {
		monthlySpending += spending
	}

	goal := countGoal(goalAddition, goalMonthlyAddition, monthlySpending, yearlyInvestmentProfitPercent)

	capitalResult := NewCapitalResult(
		startTime,
		capital,
		salary,
		monthlySpendingList,
		goal,
		yearlyInvestmentProfitPercent,
	)
	capitalResult.Append(
		[]string{
			"Дата",
			"Заработано",
			"Инвестиции",
			"Заработано всего",
			"Капитал",
			"Комментарии",
		},
	)

	calculator := accounting.Accounting{Symbol: "₽", Thousand: " ", Format: "%v %s"}

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
			capitalResult.Append([]string{currentTime.Format("2006"), "", "", "", "", ""})
		}

		previousYear = currentTime.Year()

		capital += profit + investmentProfit

		var row []string
		row = append(
			row,
			currentTime.Format("January"),
			calculator.FormatMoney(profit),
			calculator.FormatMoney(investmentProfit),
			calculator.FormatMoney(overallProfit),
			calculator.FormatMoney(capital),
			"",
		)

		// vacation
		for _, vacationMonth := range vacationMonths {
			if vacationMonth == currentTime.Month() {
				vacationPrice := int(float64(salary) * (float64(vacationSpendingPercent) / 100))
				profit -= vacationPrice
				overallProfit -= vacationPrice
				capital -= vacationPrice
				row[1] = calculator.FormatMoney(profit)
				row[3] = calculator.FormatMoney(overallProfit)
				row[4] = calculator.FormatMoney(capital)
				if row[5] != "" {
					row[5] += " | "
				}
				row[5] += fmt.Sprintf("Отпуск: %s", calculator.FormatMoney(vacationPrice))
			}
		}

		// salary growth
		for _, salaryGrowMonth := range salaryGrowMonths {
			if salaryGrowMonth == currentTime.Month() {
				salaryGrow := salary * salaryGrowPercent / 100
				salary += salaryGrow
				if row[5] != "" {
					row[5] += " | "
				}
				row[5] += fmt.Sprintf(
					"Повышение зарплаты: %s (+%s)",
					calculator.FormatMoney(salary),
					calculator.FormatMoney(salaryGrow),
				)
			}
		}

		// salary bonus
		for _, salaryBonusMonth := range salaryBonusMonths {
			if currentTime.Month() == salaryBonusMonth {
				bonus := salary * salaryBonusPercent / 100 * 6
				profit += bonus
				overallProfit += bonus
				row[1] = calculator.FormatMoney(profit)
				row[3] = calculator.FormatMoney(overallProfit)
				if row[5] != "" {
					row[5] += " | "
				}
				row[5] += fmt.Sprintf(
					"Премия: %s",
					calculator.FormatMoney(bonus),
				)
			}
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
