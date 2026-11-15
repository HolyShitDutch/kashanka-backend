-- name: ListCollections :many
select * from collections order by title;

-- name: GetCollectionBySlug :one
select * from collections where slug = $1;
