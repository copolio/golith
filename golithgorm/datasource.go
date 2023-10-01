package golithgorm

type Datasource struct {
	Ddl         Mode   `yaml:"ddl"`
	Url         string `yaml:"url"`
	Host        string `yaml:"host"`
	Protocol    string `yaml:"protocol"`
	Schema      string `yaml:"schema"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
	Charset     string `yaml:"charset"`
	ParseTime   bool   `yaml:"parseTime"`
	Loc         string `yaml:"loc"`
}

func DefaultDatasource() Datasource {
	return Datasource{
		Ddl:         NONE,
		Host:        "localhost:3306",
		Protocol:    "tcp",
		Schema:      "petstore",
		User:        "root",
		Password:    "test",
		MaxIdleConn: 10,
		MaxOpenConn: 10,
		Charset:     "utf8",
		ParseTime:   true,
		Loc:         "Local",
	}
}
