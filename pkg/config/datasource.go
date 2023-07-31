package config

type Datasource struct {
	Ddl         Ddl    `yaml:"ddl"`
	Url         string `yaml:"url"`
	Host        string `yaml:"host"`
	Protocol    string `yaml:"protocol"`
	Schema      string `yaml:"schema"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}

type Ddl string

const (
	CREATE Ddl = "CREATE"
	NONE   Ddl = "NONE"
)
