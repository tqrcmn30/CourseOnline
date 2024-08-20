-- name: CreateOrderCourse :one
INSERT INTO order_courses (usco_purchase_no, usco_tax, usco_subtotal, usco_patrx_no, usco_modified, usco_user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetOrderCourseByID :one
SELECT * FROM order_courses WHERE usco_id = $1;

-- name: GetAllOrderCourses :many
SELECT * FROM order_courses;

-- name: UpdateOrderCourse :one
UPDATE order_courses
SET usco_purchase_no = $2, usco_tax = $3, usco_subtotal = $4, usco_patrx_no = $5, usco_modified = $6, usco_user_id = $7
WHERE usco_id = $1
RETURNING *;

-- name: DeleteOrderCourse :exec
DELETE FROM order_courses WHERE usco_id = $1;