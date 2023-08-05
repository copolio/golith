package config

import (
	"fmt"
	"github.com/copolio/gin-bootify/pkg/database/ddl"
)

type Datasource struct {
	Ddl         ddl.DDL `yaml:"ddl"`
	Url         string  `yaml:"url"`
	Host        string  `yaml:"host"`
	Protocol    string  `yaml:"protocol"`
	Schema      string  `yaml:"schema"`
	User        string  `yaml:"user"`
	Password    string  `yaml:"password"`
	MaxIdleConn int     `yaml:"maxIdleConn"`
	MaxOpenConn int     `yaml:"maxOpenConn"`
	Charset     string  `yaml:"charset"`
	ParseTime   bool    `yaml:"parseTime"`
	Loc         string  `yaml:"loc"`
}

func (c *Datasource) GetMysqlDSN() string {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s", c.User, c.Password, c.Protocol, c.Host, c.Schema, c.Charset, c.ParseTime, c.Loc)
	return dsn
}
