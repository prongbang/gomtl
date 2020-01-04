package database

import (
	"github.com/prongbang/gomtl/pkg/database/driver"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataSource interface {
	GetMongoDB() *mongo.Database
}

type dataSource struct {
	MongoDB *mongo.Database
}

func (d *dataSource) GetMongoDB() *mongo.Database {
	return d.MongoDB
}

func NewDataSource(mongoDB driver.MongoDriver) DataSource {
	return &dataSource{
		MongoDB: mongoDB.Connect(),
	}
}