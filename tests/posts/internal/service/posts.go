package service

import (
	"posts/internal/models"
	"posts/internal/repos/memory"
)

type PostService struct {
	repo memory.PostRepository
}

func (s *PostService) Posts() []*models.Post {
	return s.repo.Posts()
}

func (s *PostService) CreatePost(post *models.Post) {
	s.repo.CreatePost(post)
}

func (s *PostService) PostById(id string) (*models.Post, error) {
	d, _ := s.repo.PostById(id)
	return d, nil
}

func NewPostService(repo memory.PostRepository) *PostService {

	return &PostService{repo: repo}

}
