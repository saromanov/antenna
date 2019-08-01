package config

// Config defines configuration for the antenna
type Config struct {
	HTTPAddress   string `yaml:"http_address"`
	InfluxAddress string `yaml:"influx_addrss"`
}
