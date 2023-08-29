package repository

import (
	"context"
	"encoding/json"
	"lusionhouse-server/app/domain"
	"lusionhouse-server/app/infrastructure/database/mysql"
	"lusionhouse-server/app/infrastructure/database/redis"
	"time"

	oriredis "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindByID(ctx context.Context, id int) (*domain.User, error)
	FindAll(ctx context.Context) ([]*domain.User, error)
	Save(ctx context.Context, user *domain.User) error
	DeleteByID(ctx context.Context, id int) error
	FindUserBySessionId(ctx context.Context, sessionId string) (*domain.User, error)
	SetUserSession(ctx context.Context, sessionId string, user *domain.User) error
	ClearUserSession(ctx context.Context, sessionId string) error
}

func NewUserRepository() UserRepository {
	return &userRepository{
		mysqlDb:  mysql.MysqlCli,
		redisCli: redis.SessionCli,
	}
}

type userRepository struct {
	mysqlDb  *gorm.DB
	redisCli *oriredis.Client
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	if err := r.mysqlDb.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	if err := r.mysqlDb.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.mysqlDb.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	return r.mysqlDb.Save(user).Error
}
func (r *userRepository) DeleteByID(ctx context.Context, id int) error {
	return r.mysqlDb.Update("deleted_at", time.Now()).Error
}
func (r *userRepository) FindUserBySessionId(ctx context.Context, sessionId string) (*domain.User, error) {
	var user domain.User
	var data []byte
	if err := r.redisCli.Get(ctx, sessionId).Scan(&data); err != nil {
		return nil, err
	}
	err := json.Unmarshal(data, &user)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal user failed")
	}
	return &user, nil
}
func (r *userRepository) SetUserSession(ctx context.Context, sessionId string, user *domain.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "json marshal user failed")
	}
	return r.redisCli.Set(ctx, sessionId, data, time.Hour*24*7).Err()
}

func (r *userRepository) ClearUserSession(ctx context.Context, sessionId string) error {
	return r.redisCli.Del(ctx, sessionId).Err()
}
