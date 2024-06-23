package config

import "net"

type Config struct {
	Logger   Logger
	Database PgSQL
	HTTP     HTTP
}

func (c *Config) HTTPAddr() string {
	return net.JoinHostPort(c.HTTP.Host, c.HTTP.Port)
}
