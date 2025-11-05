package usecases

import (
	"context"
	"errors"
	"strings"

	"messenger-module/entities"

	"github.com/google/uuid"
)

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase { return &UserUsecase{repo: repo} }

func (u *UserUsecase) Create(ctx context.Context, in entities.User) (entities.User, error) {
	in.Name = strings.TrimSpace(in.Name)
	if in.Name == "" {
		return entities.User{}, errors.New("name is required")
	}
	if strings.TrimSpace(in.APIKey) == "" {
		in.APIKey = uuid.New().String()
	}
	return u.repo.CreateUser(ctx, in)
}

func (u *UserUsecase) Get(ctx context.Context, id string) (entities.User, error) {
	return u.repo.GetUser(ctx, id)
}

func (u *UserUsecase) List(ctx context.Context) ([]entities.User, error) {
	return u.repo.ListUsers(ctx)
}

func (u *UserUsecase) Update(ctx context.Context, id string, in entities.User) (entities.User, error) {
	return u.repo.UpdateUser(ctx, id, in)
}

func (u *UserUsecase) Delete(ctx context.Context, id string) error { return u.repo.DeleteUser(ctx, id) }
