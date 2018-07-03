package main

import (
	"github.com/MarcoBramini/nutsify-server/datasource"
	"github.com/MarcoBramini/nutsify-server/model"
	"time"
	"fmt"
)

func main() {

	driver := datasource.NewMongoDriver("mongodb://localhost:27017", "nutsify")
	categoryDS := datasource.NewCategoryDatasource(driver, "categories")
	expenseDS := datasource.NewExpenseDatasource(driver, "expenses")
	receiptDS := datasource.NewReceiptDatasource(driver, "receipts")
	defer driver.Disconnect()
	driver.Connect()
	categoryDS.CreateCategory(model.Category{Name: "Concimazione", ID: "cat-1"})
	date, err := time.Parse(time.RFC822, "23 May 18 00:00 CEST")
	if err != nil{
		fmt.Printf("Parse failed")
	}
	receiptDS.CreateReceipt(model.Receipt{ID: "rec-1", ScanURL: "null", UploadDate: date.Unix()})
	expenseDS.CreateExpense(
		model.Expense{
			ID: "exp-1",
			ExpenseDescription: model.ExpenseDescription{
				Name:   "Fresatura",
				Date:   time.Now().Unix(),
				Amount: 3500,
				IsPaid: true,
			},
			Category: "cat-1",
			Receipt:  "rec-1"})

	exp := expenseDS.FindExpenseByID("exp-1")

	res := categoryDS.FindCategoryByID(exp.Category)
	rec := receiptDS.FindReceiptByID(exp.Receipt)

	fmt.Print(exp, res, rec)

	fmt.Print(time.Unix(exp.ExpenseDescription.Date, 0))

}
