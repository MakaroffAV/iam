package memory

import (
	"ozon-posts/internal/models"
	"sync"
)

type PostRepository struct {
	mu    sync.RWMutex
	posts map[string]*models.Post
}

func (r *PostRepository) Posts() []*models.Post {

	r.mu.Lock()
	defer r.mu.Unlock()

	p := []*models.Post{}
	for _, v := range r.posts {
		p = append(p, v)
	}

	return p

}

func (r *PostRepository) CreatePost(post *models.Post) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.posts[post.ID] = post

}

func (r *PostRepository) PostById(id string) (*models.Post, bool) {

	r.mu.Lock()
	defer r.mu.Unlock()

	p, pExist := r.posts[id]
	return p, pExist

}

func NewPostRepository() *PostRepository {

	return &PostRepository{
		posts: make(map[string]*models.Post),
	}

}
