package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

// Policy represents a row from 'unkey.policies'.
type Policy struct {
	ID        string          `json:"id"`         // id
	Name      sql.NullString  `json:"name"`       // name
	APIID     sql.NullString  `json:"api_id"`     // api_id
	CreatedAt time.Time       `json:"created_at"` // created_at
	UpdatedAt time.Time       `json:"updated_at"` // updated_at
	Version   string          `json:"version"`    // version
	Policy    json.RawMessage `json:"policy"`     // policy
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Policy] exists in the database.
func (p *Policy) Exists() bool {
	return p._exists
}

// Deleted returns true when the [Policy] has been marked for deletion
// from the database.
func (p *Policy) Deleted() bool {
	return p._deleted
}

// Insert inserts the [Policy] to the database.
func (p *Policy) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO unkey.policies (` +
		`id, name, api_id, created_at, updated_at, version, policy` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, p.ID, p.Name, p.APIID, p.CreatedAt, p.UpdatedAt, p.Version, p.Policy)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID, p.Name, p.APIID, p.CreatedAt, p.UpdatedAt, p.Version, p.Policy); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Update updates a [Policy] in the database.
func (p *Policy) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE unkey.policies SET ` +
		`name = ?, api_id = ?, created_at = ?, updated_at = ?, version = ?, policy = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, p.Name, p.APIID, p.CreatedAt, p.UpdatedAt, p.Version, p.Policy, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.Name, p.APIID, p.CreatedAt, p.UpdatedAt, p.Version, p.Policy, p.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Policy] to the database.
func (p *Policy) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for [Policy].
func (p *Policy) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO unkey.policies (` +
		`id, name, api_id, created_at, updated_at, version, policy` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`id = VALUES(id), name = VALUES(name), api_id = VALUES(api_id), created_at = VALUES(created_at), updated_at = VALUES(updated_at), version = VALUES(version), policy = VALUES(policy)`
	// run
	logf(sqlstr, p.ID, p.Name, p.APIID, p.CreatedAt, p.UpdatedAt, p.Version, p.Policy)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID, p.Name, p.APIID, p.CreatedAt, p.UpdatedAt, p.Version, p.Policy); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the [Policy] from the database.
func (p *Policy) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM unkey.policies ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// PolicyByID retrieves a row from 'unkey.policies' as a [Policy].
//
// Generated from index 'policies_id_pkey'.
func PolicyByID(ctx context.Context, db DB, id string) (*Policy, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, api_id, created_at, updated_at, version, policy ` +
		`FROM unkey.policies ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	p := Policy{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&p.ID, &p.Name, &p.APIID, &p.CreatedAt, &p.UpdatedAt, &p.Version, &p.Policy); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}