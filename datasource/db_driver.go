package datasource


type DBDriver interface {
	Connect() error
	Disconnect() error
	Create(collection string, item ...interface{}) error
	RetrieveOne(collection string, query interface{}, result interface{}) error
	RetrieveAll(collection string, query interface{}, results interface{}) error
	UpdateOne(collection string, filter interface{}, item interface{}) error
	Delete(collection string, filter interface{}) error
}
