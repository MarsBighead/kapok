package util

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

//Config Type Config for config.yaml
type Config struct {
	//	Type            string `yaml:"type"`
	Postgres []*Postgres `yaml:"postgres"`
}

// Parse config.yaml to data struct
func Parse(file string) (cfg *Config, err error) {
	cfg = new(Config)
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(body, cfg)
	if err != nil {
		return
	}
	return
}

// GetPostgres get postgres database driver and dsn info
func (cfg *Config) GetPostgres() (pg *Postgres) {
	if len(cfg.Postgres) == 0 {
		return
	}
	pg = cfg.Postgres[0]
	for _, db := range cfg.Postgres {
		if strings.ToLower(db.Switch) == "on" {
			pg = db
		}
	}
	return
}
