package model

type Customer struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Store     *Store `json:"store,omitempty"`
	StoreID   string `json:"storeId,omitempty"`
}
