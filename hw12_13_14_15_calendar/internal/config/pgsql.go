package config

type PgSQL struct {
	DSN      string
	InMemory bool `mapstructure:"in_memory"`
}
