package golithcore

type CrudRepository[T any, K any] interface {
	Save(db any, entity T) (T, error)
	FindAll(db any, where T) ([]T, error)
	FindById(db any, id K) (T, error)
	Delete(db any, entity T) error
}
