package data

type CrudRepository[T any, K any] interface {
	WithTransaction(transaction any)
	Save(entity T) (*T, error)
	FindById(id K) (*T, error)
	Delete(entity T) (err error)
}

type FpCrudRepository[T any, K any] interface {
	FpSave(entity T) func(arg any) (T, error)
	FpFindById(id K) func(arg any) (T, error)
	FpDelete(entity T) func(arg any) error
}
