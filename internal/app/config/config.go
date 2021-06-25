package config

import "time"

type Config struct {
	StartTime                     time.Time
	StartTimeCallback             func(string)   `long:"start" description:"Дата начала планирования YYYY-MM-DD"`
	InitialCapital                int            `long:"capital" default:"0" description:"Капитал на начало планирования"`
	InitialSalary                 int            `long:"salary" default:"15000" description:"Зарплата в месяц на начало планирования"`
	SalaryBonusPercent            int            `long:"salary-bonus-percent" default:"0" description:"Размер премии в процентах от зарплаты"`
	SalaryBonusMonths             []time.Month   `long:"salary-bonus-month" default:"0" description:"Месяц в который выплачивается премия (можно задать несколько, еще одним параметром)"`
	SalaryGrowthPercent           int            `long:"salary-growth-percent" default:"10" description:"Процент повышения зарплаты"`
	SalaryGrowthMonths            []time.Month   `long:"salary-growth-month" default:"1" description:"Месяц в который повышается зарплата (можно задать несколько, еще одним параметром)"`
	VacationMonths                []time.Month   `long:"vacation-month" default:"12" description:"Месяц в который планируется отпуск (можно задать несколько, еще одним параметром)"`
	VacationSpendingPercent       int            `long:"vacation-percent" default:"100" description:"Траты на отпуск в процентах от зарплаты"`
	YearlyInvestmentProfitPercent int            `long:"investment-percent" default:"10" description:"Годовой доход от инвестиций в процентах"`
	GoalAddition                  int            `long:"goal-addition" default:"0" description:"Дополнительная цель в денежном эквиваленте"`
	GoalMonthlyAddition           int            `long:"goal-monthly-addition" default:"0" description:"Дополнительная ежемесячная цель в денежном эквиваленте"`
	MonthlySpending               map[string]int `long:"monthly-spending" default:"еда:10000" description:"Месячные траты (можно задать несколько, еще одним параметром)"`
	Verbose                       bool           `long:"verbose" description:"Выводить отладочную информацию"`
}
