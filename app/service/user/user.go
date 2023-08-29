package user

import (
	"context"
	"log"
	"lusionhouse-server/app/domain"
	"lusionhouse-server/app/infrastructure/repository"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type UserServiceIml struct {
	userRepo repository.UserRepository
}

var ErrAuthFailed = errors.New("invalid credentials")

func NewUserService(userRepo repository.UserRepository) *UserServiceIml {
	return &UserServiceIml{userRepo: userRepo}
}

func (s *UserServiceIml) Authenticate(ctx context.Context, username, password string) (*domain.User, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err == repository.ErrNotFound {
		return nil, ErrAuthFailed
	}
	if err != nil {
		log.Printf("find user by username failed:%s\n", err)
		return nil, err
	}

	if user.Password != password {
		return nil, ErrAuthFailed
	}
	return user, nil
}
func (s *UserServiceIml) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	return s.userRepo.FindByUsername(ctx, username)
}

func (s *UserServiceIml) CreateUser(ctx context.Context, user *domain.User) error {
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.userRepo.Save(ctx, user)
}

func (s *UserServiceIml) DeleteUserByID(ctx context.Context, id int) error {
	return s.userRepo.DeleteByID(ctx, id)
}

func (s *UserServiceIml) FindUserBySessionId(ctx context.Context, sessionId string) (*domain.User, error) {
	user, err := s.userRepo.FindUserBySessionId(ctx, sessionId)
	if err == repository.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "find user by sessionId failed")
	}
	return user, nil
}

func (s *UserServiceIml) SetUserSession(ctx context.Context, sessionId string, user *domain.User) error {
	return s.userRepo.SetUserSession(ctx, sessionId, user)
}

func (s *UserServiceIml) ClearUserSession(ctx context.Context, sessionId string) error {
	return s.userRepo.ClearUserSession(ctx, sessionId)
}
