package repository

import (
	"context"

	"platzi.com/go/rest-ws/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostByID(ctx context.Context, id string) (*models.Post, error)
	DeletePost(ctx context.Context, id string, userId string) error
	UpdatePost(ctx context.Context, post *models.Post, userId string) error
	ListPost(ctx context.Context, page uint64) ([]*models.Post, error)
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserByID(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	return implementation.GetPostByID(ctx, id)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implementation.DeletePost(ctx, id, userId)
}

func UpdatePost(ctx context.Context, post *models.Post, userId string) error {
	return implementation.UpdatePost(ctx, post, userId)
}

func ListPost(ctx context.Context, page uint64) ([]*models.Post, error) {
	return implementation.ListPost(ctx, page)
}

func Close() error {
	return implementation.Close()
}
