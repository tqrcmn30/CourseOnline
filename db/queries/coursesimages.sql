-- name: CreateCourseImage :one
INSERT INTO courses_images (coim_filename, coim_default, coim_reme_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCourseImageByID :one
SELECT * FROM courses_images WHERE coim_id = $1;

-- name: GetAllCourseImages :many
SELECT * FROM courses_images;

-- name: UpdateCourseImage :one
UPDATE courses_images
SET coim_filename = $2, coim_default = $3, coim_reme_id = $4
WHERE coim_id = $1
RETURNING *;

-- name: DeleteCourseImage :exec
DELETE FROM courses_images WHERE coim_id = $1;