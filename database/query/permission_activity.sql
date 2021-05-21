-- name: CreatePermissionActivity :one
INSERT INTO permission_activity (
    permission_id, activity_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdatePermissionActivity :one
UPDATE permission_activity
SET activity_id = $2
WHERE permission_id = $1
    RETURNING *;

-- name: DeletePermissionActivity :exec
UPDATE permission_activity
SET deleted_at = now()
WHERE permission_id = $1;


-- name: GetPermissionActivityByPermissionId :many
SELECT * FROM permission_activity WHERE permission_id = $1 AND deleted_at IS NULL;