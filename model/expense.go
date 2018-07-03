package model

type Expense struct {
	ID   string `json:"ID" bson:"_id"`
	ExpenseDescription ExpenseDescription `json:"expenseDescription" bson:"expenseDescription"`
	Category string `json:"category" bson:"category"`
	Receipt string `json:"receipt" bson:"receipt"`
}

type ExpenseDescription struct {
	Name string	`json:"name" bson:"name"`
	Date int64	`json:"date" bson:"date"`
	Amount int	`json:"amount" bson:"amount"`
	IsPaid bool	`json:"isPaid" bson:"isPaid"`
}
