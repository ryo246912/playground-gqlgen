package graph

import "database/sql"

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
