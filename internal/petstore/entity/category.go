package entity

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Category struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func (c Category) GormDataType() string {
	return fmt.Sprintf("%d-%s", c.Id, c.Name)
}

func (c Category) GormDBDataType(*gorm.DB, *schema.Field) string {
	return "text"
}
