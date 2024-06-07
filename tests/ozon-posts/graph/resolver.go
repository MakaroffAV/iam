package graph

import (
	"ozon-posts/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostService    *service.PostService
	CommentService *service.CommentService
}
