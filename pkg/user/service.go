package user

import (
	"context"
	"fmt"
)

type UserService interface {
	CreateUser(ctx context.Context, user NewUser) error
	GetUsers(ctx context.Context) ([]NewUser, error)
}

type userService struct {
	repo UserRepository
}

func NewService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user NewUser) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *userService) GetUsers(ctx context.Context) ([]NewUser, error) {
	var users []NewUser
	cursor, err := s.repo.GetUsers(ctx)
	if err != nil{
		return nil, err
	}
	for cursor.Next(ctx) {
        var elem NewUser
        if err := cursor.Decode(&elem); err != nil {
            return nil, err
        }
        users = append(users, elem)
    }
    if err := cursor.Err(); err != nil {
        return nil, err
    }

    // Вывод результатов
    for _, task := range users {
        fmt.Printf("%+v\n", task)
    }
	return users, nil
}
