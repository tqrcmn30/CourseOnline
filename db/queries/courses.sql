-- name: CreateCourse :one
INSERT INTO courses (cours_name, cours_desc, cours_author, cours_price, cours_modified, cours_cate_id)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5)
RETURNING *;

-- name: GetCourseByID :one
SELECT * FROM courses WHERE cours_id = $1;


-- name: GetAllCourses :many
SELECT * FROM courses
JOIN course_images on cours_id = coim_id;

-- name: GetAllCoursesPaging :many
SELECT cours_name, cours_desc, cours_author, cours_price, cours_modified, cours_cate_id
FROM courses
ORDER BY cours_id
LIMIT $1 OFFSET $2;

-- name: UpdateCourse :one
UPDATE courses
SET cours_name = $2, cours_desc = $3, cours_author = $4, cours_price = $5, cours_modified = CURRENT_TIMESTAMP, cours_cate_id = $6
WHERE cours_id = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses WHERE cours_id = $1;