package database

// Config is a model for config mongodb
type Config struct {
	User         string
	Pass         string
	Host         string
	DatabaseName string
}
