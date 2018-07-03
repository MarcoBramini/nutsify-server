package datasource

import (
	"github.com/MarcoBramini/nutsify-server/model"
	"log"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
)

type CategoryDatasource struct {
	driver DBDriver
	collectionName string
}

func NewCategoryDatasource(driver DBDriver, collectionName string) *CategoryDatasource {
	datasource := new(CategoryDatasource)

	datasource.driver = driver
	datasource.collectionName = collectionName

	return datasource
}

func (cds *CategoryDatasource) CreateCategory(category model.Category) model.Category {
	err := cds.driver.Create(cds.collectionName, category)
	if err != nil {
		log.Printf(fmt.Sprint("db -> CreateCategory failed with: ", err))
	}

	return category
}

func (cds *CategoryDatasource) FindCategoryByID(categoryID string) model.Category {
	query := bson.NewDocument(bson.EC.String("_id", categoryID))

	var res model.Category
	err := cds.driver.RetrieveOne(cds.collectionName, query, &res)
	if err != nil {
		log.Printf(fmt.Sprint("db -> FindCategoryByID failed with: ", err))
	}

	return res
}

func (cds *CategoryDatasource) UpdateCategory(categoryID string, category model.Category) {
	query := bson.NewDocument(bson.EC.String("_id", categoryID))

	err := cds.driver.UpdateOne(cds.collectionName, query, category)
	if err != nil {
		log.Printf(fmt.Sprint("db -> UpdateCategory failed with: ", err))
	}
}

func (cds *CategoryDatasource) DeleteCategory(categoryID string) {
	query := bson.NewDocument(bson.EC.String("_id", categoryID))

	err := cds.driver.Delete(cds.collectionName, query)
	if err != nil {
		log.Printf(fmt.Sprint("db -> DeleteCategory failed with: ", err))
	}
}
