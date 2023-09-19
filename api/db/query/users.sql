-- name: CreateOrUpdateUser :one
INSERT INTO users (
  email, name, join_year, profile_picture_url
) VALUES (
  $1, $2, $3, $4
) 
ON CONFLICT (email) DO UPDATE SET name = $2, profile_picture_url = $4
RETURNING * ;

-- name: GetUser :one
SELECT * FROM users 
WHERE email = $1
LIMIT 1;
