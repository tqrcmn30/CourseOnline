package services

import (
	db "courseonline/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseService struct {
	*db.Queries
}

// constructor
func NewCourseService(dbConn *pgxpool.Conn) *CourseService {
	return &CourseService{
		Queries: db.New(dbConn),
	}
}
