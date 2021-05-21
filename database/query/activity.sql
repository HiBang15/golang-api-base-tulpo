-- name: GetActivityByID :one
SELECT * FROM activities WHERE id = $1 AND deleted_at is null;

-- name: CreateActivity :one
INSERT INTO activities (
    url, method, url_regex
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateActivity :one
UPDATE activities
SET url = $2, method = $3, url_regex = $4
WHERE id = $1 and deleted_at is null
RETURNING *;

-- name: DeleteActivity :exec
UPDATE activities
SET deleted_at = now()
WHERE id = $1;

-- name: GetListActivity :many
SELECT * FROM activities WHERE deleted_at is null;

-- name: CheckActivityExists :one
SELECT EXISTS (SELECT * FROM activities WHERE id = $1 AND deleted_at is null);