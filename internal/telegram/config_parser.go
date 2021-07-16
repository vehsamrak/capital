package telegram

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/vehsamrak/capital/internal/app/config"
)

type State int

const (
	stateInitial State = iota
	startTime
	initialCapital
	initialSalary
	salaryBonusPercent
	salaryBonusMonths
	salaryGrowthPercent
	salaryGrowthMonths
	vacationMonths
	vacationSpendingPercent
	yearlyInvestmentProfitPercent
	goalAddition
	goalMonthlyAddition
	monthlySpending
)

type Question struct {
	State    State
	Question string
}

type Parser struct {
	UserID string
	State  State
	Config *config.Config
}

func NewParser() *Parser {
	return &Parser{State: stateInitial, Config: &config.Config{}}
}

func (p *Parser) Parse(config *config.Config, answer string) (question string, noQuestionsLeft bool, err error) {
	if p.State != stateInitial {
		err = p.fillState(config, answer)
		if err != nil {
			return "", false, err
		}
	}

	for _, question := range questions() {
		if p.State < question.State {
			continue
		}

		return question.Question, false, nil
	}

	return "", true, nil
}

func (p *Parser) fillState(config *config.Config, answer string) error {
	switch p.State {
	case startTime:
		var err error
		config.StartTime, err = time.Parse("2006-01-02", answer)
		if err != nil {
			if _, ok := err.(*time.ParseError); ok {
				return errors.Errorf("Ошибка обработки даты. Используйте формат YYYY-MM-DD")
			}

			return errors.Errorf("Ошибка обработки даты")
		}
	case initialCapital:
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		config.InitialCapital = answerInt
	case initialSalary:
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		config.InitialSalary = answerInt
	case salaryBonusPercent:
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		config.SalaryBonusPercent = answerInt
	case salaryBonusMonths:
		var months []time.Month
		for _, month := range strings.Split(answer, ",") {
			monthInt, err := strconv.Atoi(month)
			if err != nil {
				return err
			}
			months = append(months, time.Month(monthInt))
		}

		config.SalaryBonusMonths = months
	case salaryGrowthPercent:
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		config.SalaryGrowthPercent = answerInt
	case salaryGrowthMonths:
		var months []time.Month
		for _, month := range strings.Split(answer, ",") {
			monthInt, err := strconv.Atoi(month)
			if err != nil {
				return err
			}
			months = append(months, time.Month(monthInt))
		}

		config.SalaryGrowthMonths = months
	case vacationMonths:
		var months []time.Month
		for _, month := range strings.Split(answer, ",") {
			monthInt, err := strconv.Atoi(month)
			if err != nil {
				return err
			}
			months = append(months, time.Month(monthInt))
		}

		config.VacationMonths = months
	case vacationSpendingPercent:
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		config.VacationSpendingPercent = answerInt
	case yearlyInvestmentProfitPercent:
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		config.YearlyInvestmentProfitPercent = answerInt
	case goalAddition:
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		config.GoalAddition = answerInt
	case goalMonthlyAddition:
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		config.GoalMonthlyAddition = answerInt
	case monthlySpending:
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			return err
		}
		config.MonthlySpending = make(map[string]int)
		config.MonthlySpending["Затраты"] = answerInt
	}

	return nil
}

func questions() []Question {
	return []Question{
		{State: startTime, Question: "Дата начала планирования YYYY-MM-DD"},
		{State: initialCapital, Question: "Капитал на начало планирования"},
		{State: initialSalary, Question: "Зарплата в месяц на начало планирования"},
		{State: salaryBonusPercent, Question: "Размер премии в процентах от зарплаты"},
		{
			State:    salaryBonusMonths,
			Question: "Месяц в который выплачивается премия (можно задать несколько, через запятую)",
		},
		{State: salaryGrowthPercent, Question: "Процент повышения зарплаты"},
		{
			State:    salaryGrowthMonths,
			Question: "Месяц в который повышается зарплата (можно задать несколько, через запятую)",
		},
		{State: vacationMonths, Question: "Месяц в который планируется отпуск (можно задать несколько, через запятую)"},
		{State: vacationSpendingPercent, Question: "Траты на отпуск в процентах от зарплаты"},
		{State: yearlyInvestmentProfitPercent, Question: "Годовой доход от инвестиций в процентах"},
		{State: goalAddition, Question: "Дополнительная цель в денежном эквиваленте"},
		{State: goalMonthlyAddition, Question: "Дополнительная ежемесячная цель в денежном эквиваленте"},
		{State: monthlySpending, Question: "Месячные траты"},
	}
}
