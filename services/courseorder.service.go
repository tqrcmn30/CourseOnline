package services

import (
	db "courseonline/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseorderService struct {
	*db.Queries
}

// constructor
func NewCourseorderService(dbConn *pgxpool.Conn) *CourseorderService {
	return &CourseorderService{
		Queries: db.New(dbConn),
	}
}
