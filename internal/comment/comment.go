package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("Not implemented")
)

// Store - this interface defines all of the methods
// that our service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Comment - a representation of a comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new Service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving context")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}
func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
