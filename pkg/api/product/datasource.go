package product

import (
	"context"
	"github.com/prongbang/gomtl/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type DataSource interface {
	GetList(acceptLang string) []Product
}

type dataSource struct {
	DbSource database.DataSource
}

func (d *dataSource) GetList(acceptLang string) []Product {
	product := []Product{}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := d.DbSource.GetMongoDB().Collection("product")
	cur, err := collection.Find(ctx, bson.M{"locale": acceptLang})
	if err != nil {
		return []Product{}
	}
	err = cur.All(ctx, &product)
	defer cur.Close(ctx)
	if err != nil {
		return []Product{}
	}
	if err := cur.Err(); err != nil {
		return []Product{}
	}
	return product
}

func NewDataSource(dbSource database.DataSource) DataSource {
	return &dataSource{
		DbSource: dbSource,
	}
}
