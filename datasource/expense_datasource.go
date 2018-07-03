package datasource

import (
"github.com/MarcoBramini/nutsify-server/model"
"log"
"fmt"
"github.com/mongodb/mongo-go-driver/bson"
)

type ExpenseDatasource struct {
	driver DBDriver
	collectionName string
}

func NewExpenseDatasource(driver DBDriver, collectionName string) *ExpenseDatasource {
	datasource := new(ExpenseDatasource)

	datasource.driver = driver
	datasource.collectionName = collectionName

	return datasource
}

func (cds *ExpenseDatasource) CreateExpense(expense model.Expense) model.Expense {
	err := cds.driver.Create(cds.collectionName, expense)
	if err != nil {
		log.Printf(fmt.Sprint("db -> CreateExpense failed with: ", err))
	}

	return expense
}

func (cds *ExpenseDatasource) FindExpenseByID(expenseID string) model.Expense {
	query := bson.NewDocument(bson.EC.String("_id", expenseID))

	var res model.Expense
	err := cds.driver.RetrieveOne(cds.collectionName, query, &res)
	if err != nil {
		log.Printf(fmt.Sprint("db -> FindExpenseByID failed with: ", err))
	}

	return res
}

func (cds *ExpenseDatasource) UpdateExpense(expenseID string, expense model.Expense) {
	query := bson.NewDocument(bson.EC.String("_id", expenseID))

	err := cds.driver.UpdateOne(cds.collectionName, query, expense)
	if err != nil {
		log.Printf(fmt.Sprint("db -> UpdateExpense failed with: ", err))
	}
}

func (cds *ExpenseDatasource) DeleteExpense(expenseID string) {
	query := bson.NewDocument(bson.EC.String("_id", expenseID))

	err := cds.driver.Delete(cds.collectionName, query)
	if err != nil {
		log.Printf(fmt.Sprint("db -> DeleteExpense failed with: ", err))
	}
}

