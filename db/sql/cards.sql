-- name: ListCards :many
select * from cards order by id;

-- name: ListCardsForCollection :many
select c.* from cards c
  join collections_cards cc on cc.card_id = c.id
  where cc.collection_id = $1
  order by cc.position;
