package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"ozon-posts/internal/models"
)

// Children is the resolver for the children field.
func (r *commentResolver) Children(ctx context.Context, obj *models.Comment) ([]*models.Comment, error) {
	return r.CommentService.CommentChildren(obj.ID)
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, author string, commentsAllowed bool) (*models.Post, error) {
	return r.PostService.CreatePost(
		title,
		content,
		author,
		commentsAllowed,
	)
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, postID string, parentID *string, author string, content string) (*models.Comment, error) {
	c, cErr := r.CommentService.CreateComment(
		postID,
		parentID,
		author,
		content,
		r.PostService,
	)
	if cErr != nil {
		return c, cErr
	}

	if ch, ok := r.Subscribers[postID]; ok {
		go func() {
			ch <- c
		}()
	}

	return c, nil
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *models.Post) ([]*models.Comment, error) {
	return r.CommentService.Comments(
		obj.ID,
		-1,
		-1,
	)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	return r.PostService.Posts()
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	return r.PostService.Post(id)
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, postID string, limit int, offset int) ([]*models.Comment, error) {
	return r.CommentService.Comments(
		postID,
		limit,
		offset,
	)
}

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *models.Comment, error) {
	ch := make(chan *models.Comment)

	r.Subscribers[postID] = ch

	go func() {
		<-ctx.Done()
		delete(r.Subscribers, postID)
		close(ch)
	}()

	return ch, nil
}

// Comment returns CommentResolver implementation.
func (r *Resolver) Comment() CommentResolver { return &commentResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }