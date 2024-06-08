package service

import (
	"posts/internal/models"
	"posts/internal/repos/memory"
)

type CommentService struct {
	repo memory.CommentRepository
}

func (s *CommentService) CreateComment(comment *models.Comment) {
	s.repo.CreateComment(comment)
}

func (s *CommentService) GetComments(postID string, limit, offset int) []*models.Comment {
	return s.repo.GetComments(postID, limit, offset)
}

func NewCommentService(repo memory.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}
