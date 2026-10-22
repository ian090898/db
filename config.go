package db

type DbType string

const DbTypeMongoDB DbType = "mongodb"

// Config specifies the database implementation.
type Config struct {
	Type    DbType         `yaml:"type" valid:"required,in(mongodb)"`
	Mongodb *MongodbConfig `yaml:"mongodb,omitempty" valid:"optional"`
}

// MongodbConfig contains the MongoDB connection settings.
type MongodbConfig struct {
	Name string `yaml:"name" valid:"type(string),required"`
	Url  string `yaml:"url" valid:"requrl,required"`
}
