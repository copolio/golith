package config

import (
	"github.com/copolio/gin-bootify/pkg/data/ddl"
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

func DefaultDatasource() Datasource {
	return Datasource{
		Ddl:         ddl.NONE,
		Host:        "localhost:3306",
		Protocol:    "tcp",
		Schema:      "demo",
		User:        "root",
		Password:    "test",
		MaxIdleConn: 10,
		MaxOpenConn: 10,
		Charset:     "utf8",
		ParseTime:   true,
		Loc:         "Local",
	}
}
