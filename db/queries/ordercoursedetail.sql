-- name: CreateOrderCoursesDetail :one
INSERT INTO order_courses_detail (ucde_qty, ucde_price, ucde_total_price, ucde_usco_id, ucde_cours_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetOrderCoursesDetailByID :one
SELECT * FROM order_courses_detail WHERE ucde_id = $1;

-- name: GetAllOrderCoursesDetails :many
SELECT * FROM order_courses_detail;

-- name: UpdateOrderCoursesDetail :one
UPDATE order_courses_detail
SET ucde_qty = $2, ucde_price = $3, ucde_total_price = $4, ucde_usco_id = $5, ucde_cours_id = $6
WHERE ucde_id = $1
RETURNING *;

-- name: DeleteOrderCoursesDetail :exec
DELETE FROM order_courses_detail WHERE ucde_id = $1;