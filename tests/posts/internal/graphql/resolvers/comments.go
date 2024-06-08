package resolvers

import (
	"context"
	"posts/internal/models"
	"posts/internal/service"
	"time"

	"github.com/google/uuid"
)

type CommentResolver struct {
	commentService *service.CommentService
}

func (r *CommentResolver) Comments(ctx context.Context, postID string, limit, offset int) ([]*models.Comment, error) {

	return r.commentService.GetComments(postID, limit, offset), nil
}

func (r *CommentResolver) CreateComment(ctx context.Context, postID string, parentID *string, author string, content string) (*models.Comment, error) {

	c := &models.Comment{
		ID:        uuid.New().String(),
		PostID:    postID,
		ParentID:  parentID,
		Author:    author,
		Content:   content,
		CreatedAt: time.Now().Unix(),
	}
	r.commentService.CreateComment(c)
	return c, nil

}

func NewCommentResolver(commentService *service.CommentService) *CommentResolver {
	return &CommentResolver{
		commentService: commentService,
	}
}
