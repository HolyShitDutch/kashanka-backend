-- name: CountUsers :one
select count(*) from users;

-- name: InsertUser :one
insert into users (display_name_encrypted, email_encrypted, email_hash, webauthn_id) values ($1, $2, $3, $4) returning *;
