-- name: InsertWebAuthnCredential :one
insert into webauthn_credentials (user_id, rp_id, kid, credential_json) values ($1, $2, $3, $4) returning *;
