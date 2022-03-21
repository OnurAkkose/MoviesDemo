package config

type MongoDBConfig struct {
	Host          string `yaml:"host" env:"MONGODB_HOST" env-default:""`
	Port          int    `yaml:"port" env:"MONGODB_PORT" env-default:""`
	User          string `yaml:"user" env:"MONGODB_USER" env-default:""`
	AuthMechanism string `yaml:"auth_mechanism" env:"MONGODB_AUTH_MECHANISM" env-default:""`
	Password      string `yaml:"password" env:"MONGODB_PASSWORD" env-default:""`
	Database      string `yaml:"database" env:"MONGODB_DATABASE" env-default:""`
}
