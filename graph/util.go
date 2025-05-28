package graph

import (
	"database/sql"
	"time"
)

func nullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}
func stringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func nullTimeToPtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
