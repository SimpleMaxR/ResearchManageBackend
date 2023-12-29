// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: project.sql

package database

import (
	"context"
	"time"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects (projectleader, name, researchcontent, totalfunds, startdate, enddate, qualitymonitorsid, clientid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING projectid
`

type CreateProjectParams struct {
	Projectleader     int32
	Name              string
	Researchcontent   string
	Totalfunds        float64
	Startdate         string
	Enddate           string
	Qualitymonitorsid int32
	Clientid          int32
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createProject,
		arg.Projectleader,
		arg.Name,
		arg.Researchcontent,
		arg.Totalfunds,
		arg.Startdate,
		arg.Enddate,
		arg.Qualitymonitorsid,
		arg.Clientid,
	)
	var projectid int32
	err := row.Scan(&projectid)
	return projectid, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM projects WHERE projectid = $1
`

func (q *Queries) DeleteProject(ctx context.Context, projectid int32) error {
	_, err := q.db.ExecContext(ctx, deleteProject, projectid)
	return err
}

const getParterByProject = `-- name: GetParterByProject :many
SELECT partnerid, name, address, leaderid, officephone, contactname, contactphone FROM partners WHERE partnerid IN (SELECT partnerid FROM projectpartners WHERE projectid = $1) ORDER BY partnerid
`

func (q *Queries) GetParterByProject(ctx context.Context, projectid int32) ([]Partner, error) {
	rows, err := q.db.QueryContext(ctx, getParterByProject, projectid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Partner
	for rows.Next() {
		var i Partner
		if err := rows.Scan(
			&i.Partnerid,
			&i.Name,
			&i.Address,
			&i.Leaderid,
			&i.Officephone,
			&i.Contactname,
			&i.Contactphone,
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

const getProjectById = `-- name: GetProjectById :one
SELECT projectid, projectleader, name, researchcontent, totalfunds, startdate, enddate, qualitymonitorsid, clientid FROM projects WHERE projectid = $1
`

func (q *Queries) GetProjectById(ctx context.Context, projectid int32) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectById, projectid)
	var i Project
	err := row.Scan(
		&i.Projectid,
		&i.Projectleader,
		&i.Name,
		&i.Researchcontent,
		&i.Totalfunds,
		&i.Startdate,
		&i.Enddate,
		&i.Qualitymonitorsid,
		&i.Clientid,
	)
	return i, err
}

const linkProjectPartner = `-- name: LinkProjectPartner :exec

INSERT INTO projectpartners (projectid, partnerid) VALUES ($1, $2)
`

type LinkProjectPartnerParams struct {
	Projectid int32
	Partnerid int32
}

// ProjectPartner
func (q *Queries) LinkProjectPartner(ctx context.Context, arg LinkProjectPartnerParams) error {
	_, err := q.db.ExecContext(ctx, linkProjectPartner, arg.Projectid, arg.Partnerid)
	return err
}

const linkProjectResearcher = `-- name: LinkProjectResearcher :exec

INSERT INTO projectResearchers (projectid, researcherid, joindate, workload, disposablefunds) VALUES ($1, $2, $3, $4, $5)
`

type LinkProjectResearcherParams struct {
	Projectid       int32
	Researcherid    int32
	Joindate        time.Time
	Workload        float64
	Disposablefunds float64
}

// ProjectResearcher
func (q *Queries) LinkProjectResearcher(ctx context.Context, arg LinkProjectResearcherParams) error {
	_, err := q.db.ExecContext(ctx, linkProjectResearcher,
		arg.Projectid,
		arg.Researcherid,
		arg.Joindate,
		arg.Workload,
		arg.Disposablefunds,
	)
	return err
}

const listProjectAll = `-- name: ListProjectAll :many
SELECT projectid, projectleader, name, researchcontent, totalfunds, startdate, enddate, qualitymonitorsid, clientid FROM projects
`

func (q *Queries) ListProjectAll(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.Projectid,
			&i.Projectleader,
			&i.Name,
			&i.Researchcontent,
			&i.Totalfunds,
			&i.Startdate,
			&i.Enddate,
			&i.Qualitymonitorsid,
			&i.Clientid,
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

const listProjectResearcher = `-- name: ListProjectResearcher :many
SELECT researcherid, lab_id, researcher_number, name, gender, title, age, emailaddress, leader, startdate, term, researchdirection FROM researchers WHERE researcherid IN (SELECT researcherid FROM projectResearchers WHERE projectid = $1)
`

func (q *Queries) ListProjectResearcher(ctx context.Context, projectid int32) ([]Researcher, error) {
	rows, err := q.db.QueryContext(ctx, listProjectResearcher, projectid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Researcher
	for rows.Next() {
		var i Researcher
		if err := rows.Scan(
			&i.Researcherid,
			&i.LabID,
			&i.ResearcherNumber,
			&i.Name,
			&i.Gender,
			&i.Title,
			&i.Age,
			&i.Emailaddress,
			&i.Leader,
			&i.Startdate,
			&i.Term,
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

const unlinkProjectPartner = `-- name: UnlinkProjectPartner :exec
DELETE FROM projectpartners WHERE projectid = $1 AND partnerid = $2
`

type UnlinkProjectPartnerParams struct {
	Projectid int32
	Partnerid int32
}

func (q *Queries) UnlinkProjectPartner(ctx context.Context, arg UnlinkProjectPartnerParams) error {
	_, err := q.db.ExecContext(ctx, unlinkProjectPartner, arg.Projectid, arg.Partnerid)
	return err
}

const unlinkProjectResearcher = `-- name: UnlinkProjectResearcher :exec
DELETE FROM projectResearchers WHERE projectid = $1 AND researcherid = $2
`

type UnlinkProjectResearcherParams struct {
	Projectid    int32
	Researcherid int32
}

func (q *Queries) UnlinkProjectResearcher(ctx context.Context, arg UnlinkProjectResearcherParams) error {
	_, err := q.db.ExecContext(ctx, unlinkProjectResearcher, arg.Projectid, arg.Researcherid)
	return err
}

const updateProject = `-- name: UpdateProject :one
UPDATE projects SET projectleader = $2, name = $3, researchcontent = $4, totalfunds = $5, startdate = $6, enddate = $7, qualitymonitorsid = $8, clientid = $9 WHERE projectid = $1
RETURNING projectid, projectleader, name, researchcontent, totalfunds, startdate, enddate, qualitymonitorsid, clientid
`

type UpdateProjectParams struct {
	Projectid         int32
	Projectleader     int32
	Name              string
	Researchcontent   string
	Totalfunds        float64
	Startdate         string
	Enddate           string
	Qualitymonitorsid int32
	Clientid          int32
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, updateProject,
		arg.Projectid,
		arg.Projectleader,
		arg.Name,
		arg.Researchcontent,
		arg.Totalfunds,
		arg.Startdate,
		arg.Enddate,
		arg.Qualitymonitorsid,
		arg.Clientid,
	)
	var i Project
	err := row.Scan(
		&i.Projectid,
		&i.Projectleader,
		&i.Name,
		&i.Researchcontent,
		&i.Totalfunds,
		&i.Startdate,
		&i.Enddate,
		&i.Qualitymonitorsid,
		&i.Clientid,
	)
	return i, err
}
