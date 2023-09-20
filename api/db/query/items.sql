-- name: CreateOrUpdateItem :one
INSERT INTO items (
  id,
  name,
  description,
  image_binary,
  price,
  seller_email,
  created_at
) VALUES (
  $1,$2,$3,$4,$5,$6,$7
) ON CONFLICT(id)
DO UPDATE 
SET name = $2,
    description = $3,
    image_binary = $4,
    price = $5
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
