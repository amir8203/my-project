package repository

import (
	"context"
	query "my-project/src/data/db/sqlc"
)

type UserRepository interface {

	CreateUser(ctx context.Context, username, name, phone, password string) (*query.User, error)
	GetUserByUsername(ctx context.Context, username string) (*query.User, error)
	GetUserById(ctx context.Context, id int32) (*query.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*query.User, error)
	UpdateUsername(ctx context.Context, id int32, username string) error
	UpdatePassword(ctx context.Context, id int32, password string) error
	UpdateUserName(ctx context.Context, id int32, name string) error
	UpdateUserPhone(ctx context.Context, id int32, phone string) error
	DeleteUser(ctx context.Context, id int32) error
}

type userRepository struct {
	queries *query.Queries
}

func NewUserRepository(queries *query.Queries) UserRepository {
	return &userRepository{
		queries: queries,
	}
}


func (r *userRepository) CreateUser(ctx context.Context, username, name, phone, password string) (*query.User, error) {
	user, err := r.queries.CreateUser(ctx, query.CreateUserParams{
		Username: username,
		Name:     name,
		Phone:    phone,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*query.User, error) {
	user, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int32) (*query.User, error) {
	user, err := r.queries.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByPhone(ctx context.Context, phone string) (*query.User, error) {
	user, err := r.queries.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUsername(ctx context.Context, id int32, username string) error {
	return r.queries.UpdateUsername(ctx, query.UpdateUsernameParams{
		ID:       id,
		Username: username,
	})
}

func (r *userRepository) UpdatePassword(ctx context.Context, id int32, password string) error {
	return r.queries.UpdatePassword(ctx, query.UpdatePasswordParams{
		ID:       id,
		Password: password,
	})
}

func (r *userRepository) UpdateUserName(ctx context.Context, id int32, name string) error {
	return r.queries.UpdateUserName(ctx, query.UpdateUserNameParams{
		ID:   id,
		Name: name,
	})
}

func (r *userRepository) UpdateUserPhone(ctx context.Context, id int32, phone string) error {
	return r.queries.UpdateUserPhone(ctx, query.UpdateUserPhoneParams{
		ID:    id,
		Phone: phone,
	})
}

func (r *userRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}