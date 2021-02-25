package domain

// BlogRepository is an interface for blog repositories
type BlogRepository interface {
	Save(blog Blog) (map[string]interface{}, error)
	GetByID(blogID string) (*Blog, error)
	GetAll() (*[]Blog, error)
	GetByAuthor() error
	Update() error
	Delete() error
}
