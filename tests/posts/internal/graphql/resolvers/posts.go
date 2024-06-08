package resolvers

import (
	"context"
	"posts/internal/models"
	"posts/internal/service"

	"github.com/google/uuid"
)

type PostResolver struct {
	postsService    *service.PostService
	commentsService *service.CommentService
}

func (r *PostResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	return r.postsService.Posts(), nil
}

func (r *PostResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	return r.postsService.PostById(id)
}

func (r *PostResolver) CreatePost(ctx context.Context, title string, content string, author string, commentsAllowed bool) (*models.Post, error) {

	post := &models.Post{
		ID:              uuid.New().String(),
		Title:           title,
		Content:         content,
		Author:          author,
		CommentsAllowed: commentsAllowed,
	}
	r.postsService.CreatePost(post)
	return post, nil

}

func NewPostResolver(postService *service.PostService, commentService *service.CommentService) *PostResolver {
	return &PostResolver{
		postsService:    postService,
		commentsService: commentService,
	}
}
