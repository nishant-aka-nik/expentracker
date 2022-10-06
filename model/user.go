package model

type TokenUser struct {
	Name    string `json:"name,omitempty" bson:"name"`
	Email   string `json:"email,omitempty" bson:"email"`
	Picture string `json:"picture,omitempty" bson:"picture"`
}

// TODO transaction struct and deleted transaction
type UserData struct {
	Name              string             `json:"name,omitempty" bson:"name"`
	Email             string             `json:"email,omitempty" bson:"email"`
	Picture           string             `json:"picture,omitempty" bson:"picture"`
	TotalExpenditure  float64            `json:"totalexpenditure,omitempty"`
	MoneyLeft         float64            `json:"moneyleft,omitempty"`
	PerDaySpend       float64            `json:"perdayspend,omitempty"`
	Salary            float64            `json:"salary,omitempty" bson:"salary"`
	Accounts          []Account          `json:"accounts,omitempty" bson:"accounts"`
	Cards             []Card             `json:"cards,omitempty" bson:"cards"`
	OtherExpenditures []OtherExpenditure `json:"otherexpenditures,omitempty" bson:"otherexpenditures"`
	// DeletedTransactions []DeletedTransaction `json:"deletedtransactions,omitempty" bson:"deletedtransactions"`
}
