-- name: CreateCart :one
INSERT INTO carts (cart_user_id, cart_cours_id, cart_qty, cart_price, cart_modified, cart_status, cart_cart_id)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5, $6)
RETURNING *;

-- name: GetCartByID :one
SELECT * FROM carts WHERE cart_id = $1;

-- name: GetAllCarts :many
SELECT * FROM carts;

-- name: UpdateCart :one
UPDATE carts
SET cart_user_id = $2, cart_cours_id = $3, cart_qty = $4, cart_price = $5, cart_modified = CURRENT_TIMESTAMP, cart_status = $6, cart_cart_id = $7
WHERE cart_id = $1
RETURNING *;

-- name: DeleteCart :exec
DELETE FROM carts WHERE cart_id = $1;