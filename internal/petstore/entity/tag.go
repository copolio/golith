package entity

type Tag struct {
	Id    uint64
	Name  string
	PetID int
	Pet   Pet
}
