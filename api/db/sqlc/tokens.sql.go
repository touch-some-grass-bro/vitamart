// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: tokens.sql

package db

import (
	"context"
	"time"
)

const createOrUpdateGoogleTokens = `-- name: CreateOrUpdateGoogleTokens :one
INSERT INTO google_tokens (
  user_email, 
  created_at,
  refresh_token,
  access_token,
  expires_at,
  token_type
) VALUES (
  $1,$2,$3,$4,$5,$6
) 
ON CONFLICT(user_email)
DO UPDATE 
SET created_at = $2,
    refresh_token = $3,
    access_token = $4,
    expires_at = $5,
    token_type = $6
RETURNING id, user_email, created_at, refresh_token, access_token, expires_at, token_type
`

type CreateOrUpdateGoogleTokensParams struct {
	UserEmail    string    `json:"userEmail"`
	CreatedAt    time.Time `json:"createdAt"`
	RefreshToken string    `json:"refreshToken"`
	AccessToken  string    `json:"accessToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	TokenType    string    `json:"tokenType"`
}

func (q *Queries) CreateOrUpdateGoogleTokens(ctx context.Context, arg CreateOrUpdateGoogleTokensParams) (GoogleToken, error) {
	row := q.db.QueryRowContext(ctx, createOrUpdateGoogleTokens,
		arg.UserEmail,
		arg.CreatedAt,
		arg.RefreshToken,
		arg.AccessToken,
		arg.ExpiresAt,
		arg.TokenType,
	)
	var i GoogleToken
	err := row.Scan(
		&i.ID,
		&i.UserEmail,
		&i.CreatedAt,
		&i.RefreshToken,
		&i.AccessToken,
		&i.ExpiresAt,
		&i.TokenType,
	)
	return i, err
}

const getGoogleToken = `-- name: GetGoogleToken :one
SELECT id, user_email, created_at, refresh_token, access_token, expires_at, token_type FROM google_tokens
WHERE user_email = $1
LIMIT 1
`

func (q *Queries) GetGoogleToken(ctx context.Context, userEmail string) (GoogleToken, error) {
	row := q.db.QueryRowContext(ctx, getGoogleToken, userEmail)
	var i GoogleToken
	err := row.Scan(
		&i.ID,
		&i.UserEmail,
		&i.CreatedAt,
		&i.RefreshToken,
		&i.AccessToken,
		&i.ExpiresAt,
		&i.TokenType,
	)
	return i, err
}
