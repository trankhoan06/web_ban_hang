package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type RoleUser int

const (
	AdminRole RoleUser = iota
	UserRole
	ShipperRole
	ModRole
)

func (r *RoleUser) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	switch string(bytes) {
	case "Admin":
		*r = AdminRole

	case "Shipper":
		*r = ShipperRole

	case "Mod":
		*r = ModRole
	default:
		*r = UserRole
	}
	return nil

}
func (r *RoleUser) String() string {
	switch *r {
	case AdminRole:
		return "Admin"
	case ShipperRole:

		return "Shipper"
	case ModRole:
		return "Mod"
	default:
		return "User"
	}
}
func (r *RoleUser) Value() (driver.Value, error) {
	if r == nil {

		return nil, nil
	}
	return r.String(), nil
}
func (r *RoleUser) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", r.String())), nil
}
