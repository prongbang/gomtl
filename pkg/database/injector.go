package database

func NewDbConnection() DataSource {
	mongoDriver := NewMongoDB()
	dbDataSource := NewDataSource(mongoDriver)
	return dbDataSource
}
