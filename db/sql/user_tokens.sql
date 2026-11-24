-- name: InsertUserToken :one
insert into user_tokens (context, token) values ($1, $2) returning *;
