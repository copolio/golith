package entity

type Order struct {
	Id       uint64
	PetId    uint64
	Quantity uint64
	ShipDate string
	Status   string
	Complete bool
}
