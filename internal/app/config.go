package app

import "time"

type Config struct {
	StartTimeCallback             func(string) `long:"start" description:"Дата начала планирования YYYY-MM-DD"`
	StartTime                     time.Time
	InitialCapital                int   `long:"capital" default:"0" description:"Капитал на начало планирования"`
	InitialSalary                 int   `long:"salary" default:"15000" description:"Зарплата в месяц на начало планирования"`
	YearlyInvestmentProfitPercent int   `long:"investment" default:"10" description:"Годовой доход от инвестиций в процентах"`
	SalaryBonusMonths             []int `long:"salary-bonus-month" default:"12" description:"Месяцы в которые выплачивается премия (можно задать несколько, еще одним параметром)"`

	// TODO[petr]: add all parameters
}
