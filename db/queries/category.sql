-- name: CreateCategory :one
INSERT INTO category (cate_name)
VALUES ($1)
RETURNING *;

-- name: GetCategoryByID :one
SELECT * FROM category WHERE cate_id = $1;

-- name: GetAllCategories :many
SELECT * FROM category;

-- name: UpdateCategory :one
UPDATE category
SET cate_name = $2
WHERE cate_id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category WHERE cate_id = $1;