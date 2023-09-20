-- name: BuyItem :exec
INSERT INTO transactions (
  item_id, buyer_email
) VALUES (
  $1, $2
);

