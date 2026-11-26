-- name: InsertUserToken :one
insert into user_tokens (context, token, user_id) values ($1, $2, $3) returning *;
