package config

type ServerConfig struct {
	port string // Port is the port that the server will listen on
}

func (s *ServerConfig) Port() string {
	return s.port
}

func LoadServerConfig(port string) *ServerConfig {
	return &ServerConfig{
		port: port,
	}
}
