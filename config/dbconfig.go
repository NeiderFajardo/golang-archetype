package config

type MongoConfig struct {
	uri        string
	database   string
	collection string
}

func (m *MongoConfig) URI() string {
	return m.uri
}

func (m *MongoConfig) Database() string {
	return m.database
}

func (m *MongoConfig) Collection() string {
	return m.collection
}

func LoadMongoConfig(uri, database, collection string) *MongoConfig {
	return &MongoConfig{
		uri:        uri,
		database:   database,
		collection: collection,
	}
}
