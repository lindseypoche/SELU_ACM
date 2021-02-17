package domain

var (
// BlogsService is a service type for all blog logic
// BlogsService blogsServiceInterface = &blogsService{}
)

type BlogService interface {
	CreateBlog(blog Blog) (map[string]interface{}, error)
	GetBlog(string) (*Blog, error)
	GetAllBlogs() (*[]Blog, error)
	UpdateBlog(int64)
	DeleteBlog(int64)
}

type blogService struct {
	blogRepo BlogRepository
}

// NewBlogService creates a new blog service
func NewBlogService(blogRepo BlogRepository) BlogService {
	return &blogService{blogRepo: blogRepo}
}

// CreateBlog saves a blog
func (s *blogService) CreateBlog(blog Blog) (map[string]interface{}, error) {

	result, err := s.blogRepo.Save(blog)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBlog gets a specified blog from the blog id
func (s *blogService) GetBlog(blogID string) (*Blog, error) {

	result, err := s.blogRepo.GetByID(blogID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllBlogs gets all blogs
func (s *blogService) GetAllBlogs() (*[]Blog, error) {

	blogs, err := s.blogRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// UpdateBlog ...
func (s *blogService) UpdateBlog(int64) {

}

// DeleteBlog ...
func (s *blogService) DeleteBlog(int64) {

}
