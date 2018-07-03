package datasource

import (
	"github.com/MarcoBramini/nutsify-server/model"
	"log"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
)

type ReceiptDatasource struct {
	driver         DBDriver
	collectionName string
}

func NewReceiptDatasource(driver DBDriver, collectionName string) *ReceiptDatasource {
	datasource := new(ReceiptDatasource)

	datasource.driver = driver
	datasource.collectionName = collectionName

	return datasource
}

func (cds *ReceiptDatasource) CreateReceipt(receipt model.Receipt) model.Receipt {
	err := cds.driver.Create(cds.collectionName, receipt)
	if err != nil {
		log.Printf(fmt.Sprint("db -> CreateReceipt failed with: ", err))
	}

	return receipt
}

func (cds *ReceiptDatasource) FindReceiptByID(expenseID string) model.Receipt {
	query := bson.NewDocument(bson.EC.String("_id", expenseID))

	var res model.Receipt
	err := cds.driver.RetrieveOne(cds.collectionName, query, &res)
	if err != nil {
		log.Printf(fmt.Sprint("db -> FindReceiptByID failed with: ", err))
	}

	return res
}

func (cds *ReceiptDatasource) UpdateReceipt(expenseID string, receipt model.Receipt) {
	query := bson.NewDocument(bson.EC.String("_id", expenseID))

	err := cds.driver.UpdateOne(cds.collectionName, query, receipt)
	if err != nil {
		log.Printf(fmt.Sprint("db -> UpdateReceipt failed with: ", err))
	}
}

func (cds *ReceiptDatasource) DeleteReceipt(expenseID string) {
	query := bson.NewDocument(bson.EC.String("_id", expenseID))

	err := cds.driver.Delete(cds.collectionName, query)
	if err != nil {
		log.Printf(fmt.Sprint("db -> DeleteReceipt failed with: ", err))
	}
}
