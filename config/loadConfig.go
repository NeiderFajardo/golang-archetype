package config

import "os"

// GetMongoConfig returns a MongoConfig struct with the default values
func GetMongoConfig() *MongoConfig {
	mongoUri := os.Getenv("MONGO_URI")
	mongoDatabase := os.Getenv("MONGO_DATABASE")
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	return LoadMongoConfig(mongoUri, mongoDatabase, mongoCollection)
}

// GetServerConfig returns a ServerConfig struct with the default values
func GetServerConfig() *ServerConfig {
	serverPort := os.Getenv("SERVER_PORT")
	return LoadServerConfig(serverPort)
}
