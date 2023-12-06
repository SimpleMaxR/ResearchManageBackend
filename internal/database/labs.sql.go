// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: labs.sql

package database

import (
	"context"
	"database/sql"
)

const createLab = `-- name: CreateLab :one
INSERT INTO Laboratories (Name, OfficeArea, Address, ResearchDirection) VALUES ($1, $2, $3, $4) 
RETURNING LabID
`

type CreateLabParams struct {
	Name              string
	Officearea        sql.NullFloat64
	Address           sql.NullString
	Researchdirection sql.NullString
}

func (q *Queries) CreateLab(ctx context.Context, arg CreateLabParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createLab,
		arg.Name,
		arg.Officearea,
		arg.Address,
		arg.Researchdirection,
	)
	var labid int32
	err := row.Scan(&labid)
	return labid, err
}

const deleteLab = `-- name: DeleteLab :one
DELETE FROM Laboratories WHERE LabID = $1
RETURNING labid, name, officearea, address, researchdirection
`

func (q *Queries) DeleteLab(ctx context.Context, labid int32) (Laboratory, error) {
	row := q.db.QueryRowContext(ctx, deleteLab, labid)
	var i Laboratory
	err := row.Scan(
		&i.Labid,
		&i.Name,
		&i.Officearea,
		&i.Address,
		&i.Researchdirection,
	)
	return i, err
}

const healthzDatabase = `-- name: HealthzDatabase :one
SELECT version()
`

func (q *Queries) HealthzDatabase(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, healthzDatabase)
	var version string
	err := row.Scan(&version)
	return version, err
}

const listLab = `-- name: ListLab :one
SELECT labid, name, officearea, address, researchdirection FROM Laboratories WHERE Name = $1
`

func (q *Queries) ListLab(ctx context.Context, name string) (Laboratory, error) {
	row := q.db.QueryRowContext(ctx, listLab, name)
	var i Laboratory
	err := row.Scan(
		&i.Labid,
		&i.Name,
		&i.Officearea,
		&i.Address,
		&i.Researchdirection,
	)
	return i, err
}

const listLabAll = `-- name: ListLabAll :many
SELECT labid, name, officearea, address, researchdirection FROM Laboratories
`

func (q *Queries) ListLabAll(ctx context.Context) ([]Laboratory, error) {
	rows, err := q.db.QueryContext(ctx, listLabAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Laboratory
	for rows.Next() {
		var i Laboratory
		if err := rows.Scan(
			&i.Labid,
			&i.Name,
			&i.Officearea,
			&i.Address,
			&i.Researchdirection,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLab = `-- name: UpdateLab :one
UPDATE Laboratories SET Name = $1, OfficeArea = $2, Address = $3, ResearchDirection = $4 WHERE LabID = $5
RETURNING labid, name, officearea, address, researchdirection
`

type UpdateLabParams struct {
	Name              string
	Officearea        sql.NullFloat64
	Address           sql.NullString
	Researchdirection sql.NullString
	Labid             int32
}

func (q *Queries) UpdateLab(ctx context.Context, arg UpdateLabParams) (Laboratory, error) {
	row := q.db.QueryRowContext(ctx, updateLab,
		arg.Name,
		arg.Officearea,
		arg.Address,
		arg.Researchdirection,
		arg.Labid,
	)
	var i Laboratory
	err := row.Scan(
		&i.Labid,
		&i.Name,
		&i.Officearea,
		&i.Address,
		&i.Researchdirection,
	)
	return i, err
}
