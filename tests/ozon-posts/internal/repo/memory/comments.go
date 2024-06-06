package memory

import (
	"ozon-posts/internal/models"
	"sync"
)

type CommentRepository struct {
	mu                 sync.RWMutex
	comments           map[string]*models.Comment
	postComments       map[string][]*models.Comment
	parentCommentIndex map[string][]*models.Comment
}

func (r *CommentRepository) CreateComment(comment *models.Comment) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.comments[comment.ID] = comment
	r.postComments[comment.PostID] = append(r.postComments[comment.PostID], comment)

	if comment.ParentID != nil {
		r.parentCommentIndex[*comment.ParentID] = append(r.parentCommentIndex[*comment.ParentID], comment)
	}

}

func (r *CommentRepository) GetComments(postID string, limit, offset int) []*models.Comment {

	r.mu.Lock()
	defer r.mu.Unlock()

	c := r.postComments[postID]

	if offset >= len(c) {
		return []*models.Comment{}
	}

	if offset+limit > len(c) {
		limit = len(c) - offset
	}

	return c[offset : offset+limit]

}

func NewCommentRepository() *CommentRepository {

	return &CommentRepository{
		comments:           make(map[string]*models.Comment),
		postComments:       make(map[string][]*models.Comment),
		parentCommentIndex: make(map[string][]*models.Comment),
	}

}
