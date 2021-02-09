package services

import (
	"github.com/lindseypoche/SELU_ACM/api/domain/blogs"
)

var (
	// BlogsService is a service type for all blog logic
	BlogsService blogsServiceInterface = &blogsService{}
)

type blogsServiceInterface interface {
	CreateBlog(blog blogs.Blog) (map[string]interface{}, error)
	GetBlog(string) (*blogs.Blog, error)
	GetAllBlogs() (*[]blogs.Blog, error)
	UpdateBlog(int64)
	DeleteBlog(int64)
}

type blogsService struct{}

// CreateBlog saves the blog by calling the Save function from the dao
// and returns a map of the blog id
func (s *blogsService) CreateBlog(blog blogs.Blog) (map[string]interface{}, error) {

	result, err := blog.Save()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBlog gets the specified blog by calling the GetByID function from the dao
// and returns the blog if successful
func (s *blogsService) GetBlog(blogID string) (*blogs.Blog, error) {

	var blog blogs.Blog

	if err := blog.GetByID(blogID); err != nil {
		return nil, err
	}
	return &blog, nil
}

// GetAllBlogs gets all the blogs by calling the GetAll function from the dao
// and returns them
func (s *blogsService) GetAllBlogs() (*[]blogs.Blog, error) {

	blogs, err := blogs.GetAll()
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// UpdateBlog ...
func (s *blogsService) UpdateBlog(int64) {

}

// DeleteBlog ...
func (s *blogsService) DeleteBlog(int64) {

}
