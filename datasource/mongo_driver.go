package datasource

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"fmt"
	"context"
	"reflect"
)

type MongoDriver struct {
	dbName string
	client *mongo.Client
}

func NewMongoDriver(uri string, dbName string) *MongoDriver {
	mongoDriver := new(MongoDriver)
	client, err := mongo.NewClient(uri)
	if err != nil {
		log.Print(fmt.Sprint("db -> Mongo failed with: ", err))
	}
	mongoDriver.client = client
	mongoDriver.dbName = dbName
	return mongoDriver
}

func (md *MongoDriver) Connect() error {
	err := md.client.Connect(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (md *MongoDriver) Disconnect() error {
	err := md.client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (md *MongoDriver) Create(collection string, item ...interface{}) error {
	_, err := md.client.Database(md.dbName).Collection(collection).InsertMany(context.Background(), item)
	if err != nil {
		return err
	}
	return nil
}

func (md *MongoDriver) RetrieveOne(collection string, query interface{}, result interface{}) error {
	err := md.client.Database(md.dbName).Collection(collection).FindOne(context.Background(), query).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (md *MongoDriver) RetrieveAll(collection string, query interface{}, results interface{}) error {
	cur, err := md.client.Database(md.dbName).Collection(collection).Find(context.Background(), query)
	if err != nil {
		return err
	}

	All(cur, results)
	return nil
}

func (md *MongoDriver) UpdateOne(collection string, filter interface{}, item interface{}) error {
	_, err := md.client.Database(md.dbName).Collection(collection).ReplaceOne(context.Background(), filter, item)
	if err != nil {
		return err
	}

	return nil
}

func (md *MongoDriver) Delete(collection string, filter interface{}) error {
	_, err := md.client.Database(md.dbName).Collection(collection).DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func All(cur mongo.Cursor, output interface{}) error {
	outputValue := reflect.ValueOf(output)
	outputSliceValue := outputValue.Elem()
	elementType := outputSliceValue.Type().Elem()

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {

		newElement := reflect.New(elementType)
		err := cur.Decode(newElement.Interface())
		if err != nil {
			return err
		}

		outputSliceValue = reflect.Append(outputSliceValue, newElement.Elem())
	}

	outputValue.Elem().Set(outputSliceValue)
	cur.Close(context.Background())
	return nil
}
