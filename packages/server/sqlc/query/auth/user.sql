-- name: GetByID :one
select *
from users
where id = $1
limit 1;

-- name: Create :one
insert into users default
values
returning id;