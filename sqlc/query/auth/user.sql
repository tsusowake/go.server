-- name: GetUserByID :one
select *
from users
where id = $1
limit 1;

-- name: CreateUser :one
insert into users default
values
returning id;