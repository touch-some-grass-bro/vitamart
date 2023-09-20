// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: items.sql

package db

import (
	"context"
	"time"
)

const createOrUpdateItem = `-- name: CreateOrUpdateItem :one
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
RETURNING id, name, description, image_binary, price, seller_email, issold, created_at
`

type CreateOrUpdateItemParams struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageBinary []byte    `json:"imageBinary"`
	Price       int32     `json:"price"`
	SellerEmail string    `json:"sellerEmail"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (q *Queries) CreateOrUpdateItem(ctx context.Context, arg CreateOrUpdateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createOrUpdateItem,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.ImageBinary,
		arg.Price,
		arg.SellerEmail,
		arg.CreatedAt,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ImageBinary,
		&i.Price,
		&i.SellerEmail,
		&i.Issold,
		&i.CreatedAt,
	)
	return i, err
}

const getItem = `-- name: GetItem :one
SELECT id, name, description, image_binary, price, seller_email, issold, created_at FROM items
WHERE id = $1
`

func (q *Queries) GetItem(ctx context.Context, id int64) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ImageBinary,
		&i.Price,
		&i.SellerEmail,
		&i.Issold,
		&i.CreatedAt,
	)
	return i, err
}

const getItems = `-- name: GetItems :many
SELECT id, name, description, image_binary, price, seller_email, issold, created_at FROM items
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type GetItemsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetItems(ctx context.Context, arg GetItemsParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItems, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageBinary,
			&i.Price,
			&i.SellerEmail,
			&i.Issold,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setToSold = `-- name: SetToSold :exec
UPDATE items SET isSold = true WHERE id = $1
`

func (q *Queries) SetToSold(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, setToSold, id)
	return err
}
