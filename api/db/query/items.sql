-- name: CreateOrUpdateItem :one
INSERT INTO items (
  name,
  description,
  image_binary,
  price,
  seller_email,
  created_at
) VALUES (
  $1,$2,$3,$4,$5,$6
) ON CONFLICT(id)
DO UPDATE 
SET name = $1,
    description = $2,
    image_binary = $3,
    price = $4
RETURNING *;

-- name: GetItem :one
SELECT * FROM items
WHERE id = $1;

-- name: GetItems :many
SELECT * FROM items
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: SetToSold :exec
UPDATE items SET isSold = true WHERE id = $1;
