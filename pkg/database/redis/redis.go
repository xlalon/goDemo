package redis

type Config struct {
	Address  string `yaml:"address"`
	Port     int64  `yaml:"port"`
	Password string `yaml:"password"`
	DB       int64  `yaml:"db_num"`
}
