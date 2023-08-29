package gorm

import (
	"fmt"
	"strconv"
	"strings"
)

func GetMysqlDSN(c Datasource) string {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s", c.User, c.Password, c.Protocol, c.Host, c.Schema, c.Charset, strings.ToTitle(strconv.FormatBool(c.ParseTime)), c.Loc)
	return dsn
}
