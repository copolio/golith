package entity

type PetStatus string

const (
	AVAILABLE PetStatus = "available"
	PENDING             = "pending"
	SOLD                = "sold"
)

type Pet struct {
	Id        uint64
	Category  Category
	Name      string
	PhotoUrls []string
	Tags      []Tag
	Status    PetStatus
}
