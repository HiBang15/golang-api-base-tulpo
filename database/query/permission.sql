-- name: CreatePermission :one
INSERT INTO permissions (
    name
) VALUES (
    $1
)
RETURNING *;

-- name: UpdatePermission :one
UPDATE permissions
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeletePermission :exec
UPDATE permissions
SET deleted_at = now()
WHERE id = $1;

-- name: GetPermission :one
SELECT * FROM permissions WHERE id = 1 AND deleted_at is null;