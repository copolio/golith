package entity

type Category struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func (c Category) GormDataType() string {
	return "json"
}
