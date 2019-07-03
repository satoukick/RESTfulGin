package config

import (
	"flag"
	"fmt"
	"github.com/satoukick/webserver/logs"
	"github.com/BurntSushi/toml"
)

var (
	// Conf is used to access all configurations
	Conf = &Config{}

	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "../config/postgres.toml", "config path")
}

// Config stores all configurations
type Config struct {
	PGEnv *PGEnv
}

// PGEnv stores PostgreSQL environmental parameters
type PGEnv struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

// Init initializes Config package
func Init() {
	if _, err := toml.DecodeFile(confPath, &Conf); err != nil {
		panic(err)
	}
	if Conf.PGEnv == nil {
		Conf.PGEnv = new(PGEnv)
	}
}

// GetPGEnvString returns concated string of current PGEnv
func (c *Config) GetPGEnvString() string {
	s := fmt.Sprintf("host=%s port=%s user=%s dbname=%s",
		c.PGEnv.Host, c.PGEnv.Port, c.PGEnv.User, c.PGEnv.DBName)
	if c.PGEnv.Password != "" {
		s += fmt.Sprintf(" password=%s", c.PGEnv.Password)
	}
	if c.PGEnv.SSLMode != "" {
		s += fmt.Sprintf(" sslmode=%s", c.PGEnv.SSLMode)
	}
	return s
}
