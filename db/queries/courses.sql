-- name: CreateCourse :one
INSERT INTO courses (cours_name, cours_desc, cours_author, cours_price, cours_modified, cours_cate_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetCourseByID :one
SELECT * FROM courses WHERE cours_id = $1;

-- name: GetAllCourses :many
SELECT * FROM courses;

-- name: UpdateCourse :one
UPDATE courses
SET cours_name = $2, cours_desc = $3, cours_author = $4, cours_price = $5, cours_modified = $6, cours_cate_id = $7
WHERE cours_id = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses WHERE cours_id = $1;