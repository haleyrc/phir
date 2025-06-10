-- name: CreatePatient :one
INSERT INTO patients (
  first_name, last_name
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetPatient :one
SELECT * FROM patients
WHERE id = ?;

-- name: ListPatients :many
SELECT * FROM patients
ORDER BY id;
