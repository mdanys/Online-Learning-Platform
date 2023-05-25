package domain

import (
	"context"
	"time"
)

type User struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	Name     string    `json:"name"`
	Role     string    `json:"role"`
	Deleted  bool      `json:"deleted,omitempty"`
	DtmCrt   time.Time `json:"dtm_crt"`
	DtmUpd   time.Time `json:"dtm_upd"`
}

type UserLogin struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	Name     string    `json:"name"`
	Role     string    `json:"role"`
	Deleted  bool      `json:"deleted,omitempty"`
	DtmCrt   time.Time `json:"dtm_crt"`
	DtmUpd   time.Time `json:"dtm_upd"`
	Token    string    `json:"token"`
}

type UserRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

// UserMySQLRepository is User repository in MySQL
type UserMySQLRepository interface {
	SelectUserLogin(ctx context.Context, req LoginRequest) (user User, err error)
	InsertUser(ctx context.Context, req UserRequest) (id int64, err error)
	SelectUserByID(ctx context.Context, id int64) (user User, err error)
	RemoveUser(ctx context.Context, id int64) (err error)
}

// UserUsecase is User usecase
type UserUsecase interface {
	GetUserLogin(ctx context.Context, req LoginRequest) (user UserLogin, err error)
	CreateUser(ctx context.Context, req UserRequest) (user User, err error)
	DeleteUser(ctx context.Context, id int64) (err error)
}