package renderer

import (
	"fmt"
	"os"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"

	"github.com/vehsamrak/capital/internal/app"
)

type Console struct {
}

func (c Console) Render(capitalResult *app.CapitalResult) (render string, err error) {
	table := tablewriter.NewWriter(os.Stdout)

	for i, row := range capitalResult.Table {
		if i == 0 {
			table.SetHeader(capitalResult.Table[0])
			continue
		}

		table.Append(row)
	}

	table.Render()

	printStatistics(capitalResult)

	return
}

func printStatistics(capitalResult *app.CapitalResult) {
	calculator := accounting.Accounting{Symbol: "₽", Thousand: " ", Format: "%v %s"}

	fmt.Printf("\n")
	fmt.Printf("Цель: %s\n", calculator.FormatMoney(capitalResult.Goal))
	fmt.Printf("Пассивный доход в месяц: %s\n", calculator.FormatMoney(capitalResult.MonthlyGoal))
	fmt.Printf("Дата начала: %s\n", capitalResult.InitialDate.Format("2.01.2006"))

	yearsToGoal, monthsToGoal, _, _, _, _ := countTimeDiff(capitalResult.InitialDate, capitalResult.GoalDate)
	fmt.Printf(
		"Дата достижения цели: %s (%d лет %d мес)\n",
		capitalResult.GoalDate.Format("January 2006"),
		yearsToGoal,
		monthsToGoal,
	)
	fmt.Printf("Начальный капитал: %s\n", calculator.FormatMoney(capitalResult.InitialCapital))
	fmt.Printf("Капитал на конец периода: %s\n", calculator.FormatMoney(capitalResult.ResultCapital))
	fmt.Printf("Начальная зарплата: %s\n", calculator.FormatMoney(capitalResult.InitialSalary))
	fmt.Printf("Зарплата на конец периода: %s\n", calculator.FormatMoney(capitalResult.ResultSalary))
	fmt.Printf("Ежемесячные траты:\n")

	var spendingSum int
	for spendingName, spendingValue := range capitalResult.MonthlySpending {
		spendingSum += spendingValue
		fmt.Printf("\t%s: %s\n", spendingName, calculator.FormatMoney(spendingValue))
	}

	fmt.Printf("\tИтого: %s\n", calculator.FormatMoney(spendingSum))

	fmt.Printf("\n\n--------------------\nДля справки запустите приложение с флагом --help.\n")
	fmt.Printf("Petr Karmashev, 2021 | TLG: @vehsamrak | https://github.com/vehsamrak/capital\n")

}

func countTimeDiff(timeFrom, timeTo time.Time) (year, month, day, hour, min, sec int) {
	if timeFrom.Location() != timeTo.Location() {
		timeTo = timeTo.In(timeFrom.Location())
	}
	if timeFrom.After(timeTo) {
		timeFrom, timeTo = timeTo, timeFrom
	}
	yearFrom, monthFrom, dayFrom := timeFrom.Date()
	yearTo, monthTo, dayTo := timeTo.Date()

	hourFrom, minuteFrom, secondFrom := timeFrom.Clock()
	hourTo, minuteTo, secondTo := timeTo.Clock()

	year = yearTo - yearFrom
	month = int(monthTo - monthFrom)
	day = dayTo - dayFrom
	hour = hourTo - hourFrom
	min = minuteTo - minuteFrom
	sec = secondTo - secondFrom

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(yearFrom, monthFrom, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
