package graph

import (
	"ozon-posts/internal/repo/memory"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostStorage    *memory.PostRepository
	CommentStorage *memory.CommentRepository
}
