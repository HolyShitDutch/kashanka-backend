-- name: CountUsers :one
select count(*) from users;

-- name: InsertUser :one
insert into users (display_name_encrypted, email_encrypted, email_hash, webauthn_id) values ($1, $2, $3, $4) returning *;

-- name: GetUserByAccessToken :one
select u.* from users u
join user_tokens ut on ut.user_id = u.id
where ut.context = 'access' and ut.token = $1 and ut.inserted_at > now() - @MaxAge::interval;
