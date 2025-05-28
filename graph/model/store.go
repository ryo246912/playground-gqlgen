package model

import "time"

type Store struct {
	ID             string    `json:"id"`
	LastUpdate     time.Time `json:"lastUpdate"`
	AddressID      string    `json:"address_id,omitempty"`
	Address        *Address  `json:"address,omitempty"`
	ManagerStaffID string    `json:"manager_staff_id,omitempty"`
	ManagerStaffs  []*Staff  `json:"staffs,omitempty"`
}
