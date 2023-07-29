package database

type CrudRepository[T any, K any] interface {
	WithTransaction(transaction any)
	Save(entity T) (*T, error)
	FindById(id K) (*T, error)
	Delete(entity T) (err error)
}
