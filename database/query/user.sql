-- name: CreateUserAccount :one
INSERT INTO user_account (
    "first_name", "last_name" , "full_name", "address", "email", "password", "phone_number", "accepts_marketing", "code_verify_email", "verify_email", "password_cost", "registration_time", "email_confirmation_token" , "user_status", "password_reminder_token", "password_reminder_expire"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
)
RETURNING *;

-- name: GetUserAccountByUsernameOrEmail :one
SELECT * FROM user_account WHERE email = $1 AND deleted_at is null LIMIT 1;

-- name: CheckEmailExists :one
SELECT EXISTS (SELECT * FROM user_account WHERE email = $1 AND deleted_at is null);

-- name: CheckPhoneNoExists :one
SELECT EXISTS (SELECT * FROM user_account WHERE phone_number = $1 AND deleted_at is null);