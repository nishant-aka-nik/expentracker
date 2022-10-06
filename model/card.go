package model

type Card struct {
	Alias        string
	Number       int64
	Expenditures Expenditure
}
type Expenditure struct {
	ThisMonthName   string
	ThisMonthAmount float64

	NextMonthName   string
	NextMonthAmount float64
}
