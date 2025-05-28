package model

import "time"

type Customer struct {
	ID         string     `json:"id"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	Active     bool       `json:"active"`
	CreateDate time.Time  `json:"create_date"`
	LastUpdate *time.Time `json:"last_update,omitempty"`
	Store      *Store     `json:"store,omitempty"`
	StoreID    string     `json:"storeId,omitempty"`
}
