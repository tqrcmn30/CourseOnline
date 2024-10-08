-- name: CreateCart :one
INSERT INTO carts (cart_user_id, cart_cours_id, cart_qty, cart_price, cart_modified, cart_status)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5)
RETURNING *;

-- name: GetCartByID :one
SELECT * FROM carts cr 
JOIN user on cr.cart_user_id=cu.cart_user_id
JOIN courses on cr.cart_cours_id=cu.cart_cours_id
WHERE cart_id = $1;

-- name: GetCartByUserID :many
SELECT * FROM carts cr 
JOIN users on user_id=cr.cart_user_id
JOIN courses on cours_id=cr.cart_cours_id
WHERE cart_user_id = $1;

-- name: GetCartByUserandCourse :one
SELECT * FROM carts
WHERE cart_user_id = $1 and cart_cours_id = $2 LIMIT 1;

-- name: GetAllCarts :many
SELECT * FROM carts;

-- name: UpdateCart :one
UPDATE carts
SET cart_qty = $1, cart_price = $2
WHERE cart_id = $3
RETURNING *;

-- name: DeleteCart :exec
DELETE FROM carts WHERE cart_id = $1;