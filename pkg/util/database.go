package util

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Postgres database configure information
type Postgres struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Switch   string `yaml:"switch"`
	Port     int    `yaml:"port"`
	Sslmode  string `yaml:"sslmode"`
}

//Conn a database or other data source connection aggregation
type Conn struct {
	DB *sql.DB
}

// OpenPgConnection open databse Postgres connection
func (pg *Postgres) OpenPgConnection(ctx context.Context) (*sql.DB, error) {
	host := pg.Host
	if pg.Port != 0 && pg.Port != 5432 {
		host = fmt.Sprintf("%s:%d", pg.Host, pg.Port)
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", pg.Username, pg.Password, host, pg.DBName, pg.Sslmode)
	return sql.Open(pg.Driver, dsn)

}
