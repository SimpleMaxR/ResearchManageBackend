// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: researcher.sql

package database

import (
	"context"
)

const createResearcher = `-- name: CreateResearcher :one
INSERT INTO Researchers (LabID, Name, Gender, Title, Age, ResearchDirection, Leader) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING ResearcherID
`

type CreateResearcherParams struct {
	Labid             int32
	Name              string
	Gender            string
	Title             string
	Age               int32
	Researchdirection string
	Leader            bool
}

func (q *Queries) CreateResearcher(ctx context.Context, arg CreateResearcherParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createResearcher,
		arg.Labid,
		arg.Name,
		arg.Gender,
		arg.Title,
		arg.Age,
		arg.Researchdirection,
		arg.Leader,
	)
	var researcherid int32
	err := row.Scan(&researcherid)
	return researcherid, err
}

const deleteResearcher = `-- name: DeleteResearcher :one
DELETE FROM Researchers WHERE ResearcherID = $1
RETURNING researcherid, labid, researchnumber, name, gender, title, age, emailaddress, researchdirection, leader
`

func (q *Queries) DeleteResearcher(ctx context.Context, researcherid int32) (Researcher, error) {
	row := q.db.QueryRowContext(ctx, deleteResearcher, researcherid)
	var i Researcher
	err := row.Scan(
		&i.Researcherid,
		&i.Labid,
		&i.Researchnumber,
		&i.Name,
		&i.Gender,
		&i.Title,
		&i.Age,
		&i.Emailaddress,
		&i.Researchdirection,
		&i.Leader,
	)
	return i, err
}

const listResearcher = `-- name: ListResearcher :one
SELECT researcherid, labid, researchnumber, name, gender, title, age, emailaddress, researchdirection, leader FROM Researchers WHERE ResearcherID = $1
`

func (q *Queries) ListResearcher(ctx context.Context, researcherid int32) (Researcher, error) {
	row := q.db.QueryRowContext(ctx, listResearcher, researcherid)
	var i Researcher
	err := row.Scan(
		&i.Researcherid,
		&i.Labid,
		&i.Researchnumber,
		&i.Name,
		&i.Gender,
		&i.Title,
		&i.Age,
		&i.Emailaddress,
		&i.Researchdirection,
		&i.Leader,
	)
	return i, err
}

const listResearcherAll = `-- name: ListResearcherAll :many
SELECT researcherid, labid, researchnumber, name, gender, title, age, emailaddress, researchdirection, leader FROM Researchers
`

func (q *Queries) ListResearcherAll(ctx context.Context) ([]Researcher, error) {
	rows, err := q.db.QueryContext(ctx, listResearcherAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Researcher
	for rows.Next() {
		var i Researcher
		if err := rows.Scan(
			&i.Researcherid,
			&i.Labid,
			&i.Researchnumber,
			&i.Name,
			&i.Gender,
			&i.Title,
			&i.Age,
			&i.Emailaddress,
			&i.Researchdirection,
			&i.Leader,
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

const listResearcherByLab = `-- name: ListResearcherByLab :many
SELECT researcherid, labid, researchnumber, name, gender, title, age, emailaddress, researchdirection, leader FROM Researchers WHERE LabID = $1
`

func (q *Queries) ListResearcherByLab(ctx context.Context, labid int32) ([]Researcher, error) {
	rows, err := q.db.QueryContext(ctx, listResearcherByLab, labid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Researcher
	for rows.Next() {
		var i Researcher
		if err := rows.Scan(
			&i.Researcherid,
			&i.Labid,
			&i.Researchnumber,
			&i.Name,
			&i.Gender,
			&i.Title,
			&i.Age,
			&i.Emailaddress,
			&i.Researchdirection,
			&i.Leader,
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

const updateResearcher = `-- name: UpdateResearcher :one
UPDATE Researchers SET LabID = $1, Name = $2, Gender = $3, Title = $4, Age = $5, ResearchDirection = $6, Leader = $7 WHERE ResearcherID = $8
RETURNING researcherid, labid, researchnumber, name, gender, title, age, emailaddress, researchdirection, leader
`

type UpdateResearcherParams struct {
	Labid             int32
	Name              string
	Gender            string
	Title             string
	Age               int32
	Researchdirection string
	Leader            bool
	Researcherid      int32
}

func (q *Queries) UpdateResearcher(ctx context.Context, arg UpdateResearcherParams) (Researcher, error) {
	row := q.db.QueryRowContext(ctx, updateResearcher,
		arg.Labid,
		arg.Name,
		arg.Gender,
		arg.Title,
		arg.Age,
		arg.Researchdirection,
		arg.Leader,
		arg.Researcherid,
	)
	var i Researcher
	err := row.Scan(
		&i.Researcherid,
		&i.Labid,
		&i.Researchnumber,
		&i.Name,
		&i.Gender,
		&i.Title,
		&i.Age,
		&i.Emailaddress,
		&i.Researchdirection,
		&i.Leader,
	)
	return i, err
}