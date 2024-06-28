package config

func GetMongoConfig() *MongoConfig {
	return LoadMongoConfig("mongodb://root:password@mongodb:27017/", "test", "products")
}
