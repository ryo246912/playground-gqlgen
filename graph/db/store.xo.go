package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// Store represents a row from 'sakila.store'.
type Store struct {
	bun.BaseModel `bun:"table:store"`

	StoreID        uint8     `json:"store_id"`         // store_id
	ManagerStaffID uint8     `json:"manager_staff_id"` // manager_staff_id
	AddressID      uint16    `json:"address_id"`       // address_id
	LastUpdate     time.Time `json:"last_update"`      // last_update
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Store] exists in the database.
func (s *Store) Exists() bool {
	return s._exists
}

// Deleted returns true when the [Store] has been marked for deletion
// from the database.
func (s *Store) Deleted() bool {
	return s._deleted
}

// Insert inserts the [Store] to the database.
func (s *Store) Insert(ctx context.Context, db DB) error {
	switch {
	case s._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case s._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO sakila.store (` +
		`manager_staff_id, address_id, last_update` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`
	// run
	logf(sqlstr, s.ManagerStaffID, s.AddressID, s.LastUpdate)
	res, err := db.ExecContext(ctx, sqlstr, s.ManagerStaffID, s.AddressID, s.LastUpdate)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	s.StoreID = uint8(id)
	// set exists
	s._exists = true
	return nil
}

// Update updates a [Store] in the database.
func (s *Store) Update(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case s._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE sakila.store SET ` +
		`manager_staff_id = ?, address_id = ?, last_update = ? ` +
		`WHERE store_id = ?`
	// run
	logf(sqlstr, s.ManagerStaffID, s.AddressID, s.LastUpdate, s.StoreID)
	if _, err := db.ExecContext(ctx, sqlstr, s.ManagerStaffID, s.AddressID, s.LastUpdate, s.StoreID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Store] to the database.
func (s *Store) Save(ctx context.Context, db DB) error {
	if s.Exists() {
		return s.Update(ctx, db)
	}
	return s.Insert(ctx, db)
}

// Upsert performs an upsert for [Store].
func (s *Store) Upsert(ctx context.Context, db DB) error {
	switch {
	case s._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO sakila.store (` +
		`store_id, manager_staff_id, address_id, last_update` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`manager_staff_id = VALUES(manager_staff_id), address_id = VALUES(address_id), last_update = VALUES(last_update)`
	// run
	logf(sqlstr, s.StoreID, s.ManagerStaffID, s.AddressID, s.LastUpdate)
	if _, err := db.ExecContext(ctx, sqlstr, s.StoreID, s.ManagerStaffID, s.AddressID, s.LastUpdate); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Delete deletes the [Store] from the database.
func (s *Store) Delete(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return nil
	case s._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM sakila.store ` +
		`WHERE store_id = ?`
	// run
	logf(sqlstr, s.StoreID)
	if _, err := db.ExecContext(ctx, sqlstr, s.StoreID); err != nil {
		return logerror(err)
	}
	// set deleted
	s._deleted = true
	return nil
}

// StoreByAddressID retrieves a row from 'sakila.store' as a [Store].
//
// Generated from index 'idx_fk_address_id'.
func StoreByAddressID(ctx context.Context, db DB, addressID uint16) ([]*Store, error) {
	// query
	const sqlstr = `SELECT ` +
		`store_id, manager_staff_id, address_id, last_update ` +
		`FROM sakila.store ` +
		`WHERE address_id = ?`
	// run
	logf(sqlstr, addressID)
	rows, err := db.QueryContext(ctx, sqlstr, addressID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Store
	for rows.Next() {
		s := Store{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&s.StoreID, &s.ManagerStaffID, &s.AddressID, &s.LastUpdate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// StoreByManagerStaffID retrieves a row from 'sakila.store' as a [Store].
//
// Generated from index 'idx_unique_manager'.
func StoreByManagerStaffID(ctx context.Context, db DB, managerStaffID uint8) (*Store, error) {
	// query
	const sqlstr = `SELECT ` +
		`store_id, manager_staff_id, address_id, last_update ` +
		`FROM sakila.store ` +
		`WHERE manager_staff_id = ?`
	// run
	logf(sqlstr, managerStaffID)
	s := Store{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, managerStaffID).Scan(&s.StoreID, &s.ManagerStaffID, &s.AddressID, &s.LastUpdate); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}

// StoreByStoreID retrieves a row from 'sakila.store' as a [Store].
//
// Generated from index 'store_store_id_pkey'.
func StoreByStoreID(ctx context.Context, db DB, storeID uint8) (*Store, error) {
	// query
	const sqlstr = `SELECT ` +
		`store_id, manager_staff_id, address_id, last_update ` +
		`FROM sakila.store ` +
		`WHERE store_id = ?`
	// run
	logf(sqlstr, storeID)
	s := Store{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, storeID).Scan(&s.StoreID, &s.ManagerStaffID, &s.AddressID, &s.LastUpdate); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}

// Address returns the Address associated with the [Store]'s (AddressID).
//
// Generated from foreign key 'fk_store_address'.
func (s *Store) Address(ctx context.Context, db DB) (*Address, error) {
	return AddressByAddressID(ctx, db, s.AddressID)
}

// Staff returns the Staff associated with the [Store]'s (ManagerStaffID).
//
// Generated from foreign key 'fk_store_staff'.
func (s *Store) Staff(ctx context.Context, db DB) (*Staff, error) {
	return StaffByStaffID(ctx, db, s.ManagerStaffID)
}
