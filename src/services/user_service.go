package services

import (
	"context"
	"fmt"
	"my-project/src/api/dto"
	"my-project/src/common"
	"my-project/src/config"
	"my-project/src/data/db"
	query "my-project/src/data/db/sqlc"
	repository "my-project/src/data/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	cfg        *config.Config
	tokenService *TokenService
	db *pgxpool.Pool
	query *query.Queries
	repository repository.UserRepository
}

func NewUserService(cfg *config.Config) *UserService {
	db := db.GetDb()
	query := query.New(db)
	repository := repository.NewUserRepository(query)
	return &UserService{
		cfg: cfg,
		tokenService: NewTokenService(cfg),
		db: db,
		query: query,
		repository: repository,
	}
}


func (s *UserService) LoginByUsername(username, password string) (*dto.TokenDetail, error) {
	user, _ := s.repository.GetUserByUsername(context.Background(), username)
	if user == nil {
		return nil, fmt.Errorf("user not exist")
	}
	hashedPassword := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	tokenDto := &TokenDto{UserId: int(user.ID), Username: user.Username, MobileNumber: user.Phone}

	token, err := s.tokenService.GenerateToken(tokenDto)

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *UserService) RegisterByUsername(req dto.RegisterUserByUsernameRequest) error {

	exists, _ := s.repository.GetUserByUsername(context.Background() ,req.Username)
	if exists != nil {
		return fmt.Errorf("username exist")
	}

	exists, _ = s.repository.GetUserByPhone(context.Background(), req.Username)
	if exists != nil {
		return fmt.Errorf("phone exist")
	}

	hashedPassword, err := common.HashPassword(req.Password)
	if err != nil {
		return err
	}

	_, err = s.repository.CreateUser(context.Background(), req.Username, req.Name, req.Phone, hashedPassword)
	return err
}

func (s *UserService) GetInfo(id int) (*dto.UserProfileResponse ,error) {
	// nokte ine ke context bayad biad sotoh payin va id inja ham mitone estekhraj she
	res := new(dto.UserProfileResponse)
	user, err := s.repository.GetUserById(context.Background(), int32(id))
	if err != nil {
		return nil, fmt.Errorf("user not exist")
	}

	res.ID = int64(user.ID)
	res.Name = user.Name
	res.Phone = user.Phone
	res.Username = user.Username

	return res, nil
}


func (s *UserService) UpdateUserProfile(req dto.UpdateUserProfileRequest, id int) error {

	user, _ := s.repository.GetUserByPhone(context.Background(), req.Phone)
		if user != nil {
			return fmt.Errorf("phone exist")
		}
		user, _ = s.repository.GetUserByUsername(context.Background(), req.Username)
		if user != nil {
			return fmt.Errorf("username exist")
		}

	if req.Username != "" {
		s.repository.UpdateUsername(context.Background(), int32(id), req.Username)
	}

	if req.Phone != "" {
		s.repository.UpdateUserPhone(context.Background(), int32(id), req.Phone)
	}

	if req.Name != "" {
		s.repository.UpdateUserName(context.Background(), int32(id), req.Name)
	}

	if req.Password != "" {
		hashedPass, _:= common.HashPassword(req.Password)
		s.repository.UpdatePassword(context.Background(), int32(id), hashedPass)
	}

	return nil
}


func (s *UserService) DeleteAccount(id int) error {
	err := s.repository.DeleteUser(context.Background(), int32(id))

	if err != nil {
		return fmt.Errorf("can't delete")
	}

	return nil
}
