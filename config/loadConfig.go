package config

// GetMongoConfig returns a MongoConfig struct with the default values
func GetMongoConfig() *MongoConfig {
	return LoadMongoConfig("mongodb://root:password@mongodb:27017/", "test", "products")
}
