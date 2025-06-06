package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// City represents a row from 'sakila.city'.
type City struct {
	bun.BaseModel `bun:"table:city"`

	CityID     uint16    `json:"city_id"`     // city_id
	City       string    `json:"city"`        // city
	CountryID  uint16    `json:"country_id"`  // country_id
	LastUpdate time.Time `json:"last_update"` // last_update
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [City] exists in the database.
func (c *City) Exists() bool {
	return c._exists
}

// Deleted returns true when the [City] has been marked for deletion
// from the database.
func (c *City) Deleted() bool {
	return c._deleted
}

// Insert inserts the [City] to the database.
func (c *City) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO sakila.city (` +
		`city, country_id, last_update` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`
	// run
	logf(sqlstr, c.City, c.CountryID, c.LastUpdate)
	res, err := db.ExecContext(ctx, sqlstr, c.City, c.CountryID, c.LastUpdate)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	c.CityID = uint16(id)
	// set exists
	c._exists = true
	return nil
}

// Update updates a [City] in the database.
func (c *City) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE sakila.city SET ` +
		`city = ?, country_id = ?, last_update = ? ` +
		`WHERE city_id = ?`
	// run
	logf(sqlstr, c.City, c.CountryID, c.LastUpdate, c.CityID)
	if _, err := db.ExecContext(ctx, sqlstr, c.City, c.CountryID, c.LastUpdate, c.CityID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [City] to the database.
func (c *City) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Upsert performs an upsert for [City].
func (c *City) Upsert(ctx context.Context, db DB) error {
	switch {
	case c._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO sakila.city (` +
		`city_id, city, country_id, last_update` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`city = VALUES(city), country_id = VALUES(country_id), last_update = VALUES(last_update)`
	// run
	logf(sqlstr, c.CityID, c.City, c.CountryID, c.LastUpdate)
	if _, err := db.ExecContext(ctx, sqlstr, c.CityID, c.City, c.CountryID, c.LastUpdate); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Delete deletes the [City] from the database.
func (c *City) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM sakila.city ` +
		`WHERE city_id = ?`
	// run
	logf(sqlstr, c.CityID)
	if _, err := db.ExecContext(ctx, sqlstr, c.CityID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CityByCityID retrieves a row from 'sakila.city' as a [City].
//
// Generated from index 'city_city_id_pkey'.
func CityByCityID(ctx context.Context, db DB, cityID uint16) (*City, error) {
	// query
	const sqlstr = `SELECT ` +
		`city_id, city, country_id, last_update ` +
		`FROM sakila.city ` +
		`WHERE city_id = ?`
	// run
	logf(sqlstr, cityID)
	c := City{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, cityID).Scan(&c.CityID, &c.City, &c.CountryID, &c.LastUpdate); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}

// CityByCountryID retrieves a row from 'sakila.city' as a [City].
//
// Generated from index 'idx_fk_country_id'.
func CityByCountryID(ctx context.Context, db DB, countryID uint16) ([]*City, error) {
	// query
	const sqlstr = `SELECT ` +
		`city_id, city, country_id, last_update ` +
		`FROM sakila.city ` +
		`WHERE country_id = ?`
	// run
	logf(sqlstr, countryID)
	rows, err := db.QueryContext(ctx, sqlstr, countryID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*City
	for rows.Next() {
		c := City{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&c.CityID, &c.City, &c.CountryID, &c.LastUpdate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Country returns the Country associated with the [City]'s (CountryID).
//
// Generated from foreign key 'fk_city_country'.
func (c *City) Country(ctx context.Context, db DB) (*Country, error) {
	return CountryByCountryID(ctx, db, c.CountryID)
}
