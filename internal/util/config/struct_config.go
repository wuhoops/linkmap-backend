package config

type config struct {
	Environment environment `yaml:"environment" validate:"gte=1,lte=2"`
	LogLevel    uint32      `yaml:"log_level" validate:"required"`
	Address     string      `yaml:"address" validate:"required"`
	Cors        []string    `yaml:"cors" validate:"required"`
	DBConn      string      `yaml:"DBConn" validate:"required"`
	Secret      string      `yaml:"secret" validate:"required"`
	Redis       string      `yaml:"redis" validate:"required"`
}
