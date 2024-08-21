package services

import (
	"context"
	db "courseonline/db/sqlc"
	"courseonline/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	db.Querier
	Signup(ctx context.Context, userReq models.CreateUserReq) (*models.UserResponse, *models.Error)
	Signin(ctx context.Context, userReq models.CreateUserReq) (*models.UserResponse, *models.Error)
	Signout(ctx context.Context, accessToken string) *models.Error
}

type StoreManager struct {
	*db.Queries // implements Querier
	dbConn      *pgxpool.Conn
}

func NewStoreManager(dbConn *pgxpool.Conn) Store {
	return &StoreManager{
		Queries: db.New(dbConn),
		dbConn:  dbConn,
	}
}
