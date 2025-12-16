package repo

import (
	"context"
	"errors"

	"strings"

	"github.com/MatveySotnikov/Go_bcrypt/internal/core"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")
var ErrEmailTaken = errors.New("email already in use")

type UserRepo struct{ db *gorm.DB }

func NewUserRepo(db *gorm.DB) *UserRepo { return &UserRepo{db: db} }

func (r *UserRepo) AutoMigrate() error {
	return r.db.AutoMigrate(&core.User{})
}

func (r *UserRepo) Create(ctx context.Context, u *core.User) error {
	err := r.db.WithContext(ctx).Create(u).Error

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrEmailTaken
			}
		}

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") ||
			strings.Contains(err.Error(), "23505") {
			return ErrEmailTaken
		}

		// Если это другая ошибка БД, возвращаем ее (что приведет к 500)
		return err
	}
	return nil
}

func (r *UserRepo) ByEmail(ctx context.Context, email string) (core.User, error) {
	var u core.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return core.User{}, ErrUserNotFound
	}
	return u, err
}
