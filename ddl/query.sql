-- name: CreatePatient :one
INSERT INTO patients (
  first_name, last_name
) VALUES (
  ?, ?
)
RETURNING *;

-- name: ListPatients :many
SELECT * FROM patients
ORDER BY id;
