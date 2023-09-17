package entity

type PetStatus string

const (
	AVAILABLE PetStatus = "available"
	PENDING   PetStatus = "pending"
	SOLD      PetStatus = "sold"
)

type Pet struct {
	Id        uint64
	Category  Category
	Name      string
	PhotoUrls []string
	Tags      []Tag
	Status    PetStatus
}
