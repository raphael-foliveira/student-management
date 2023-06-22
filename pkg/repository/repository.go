package repository

type Repository[T interface{}] interface {
	All() ([]T, error)
	Find(id uint) (*T, error)
	Create(*T) error
	Delete(id uint) error
}
