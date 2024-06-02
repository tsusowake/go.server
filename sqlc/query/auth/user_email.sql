-- name: GetUserEmailByUserID :one
select user_id, email
from user_emails
where user_id = $1
limit 1;
