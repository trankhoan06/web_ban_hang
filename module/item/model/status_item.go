package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"main.go/common"
	"strings"
)

type StatusItem int

const (
	StatusItemDoing StatusItem = iota
	StatusItemDone
	StatusItemDeleted
)

var AllStatusItem = [3]string{"Doing", "Done", "Deleted"}

func (item *StatusItem) String() string {
	return AllStatusItem[*item]
}
func PathStatusItem(s string) (StatusItem, error) {
	for item := range AllStatusItem {
		if s == AllStatusItem[item] {
			return StatusItem(item), nil
		}
	}
	return StatusItem(0), common.ErrFoundItem

}
func (item *StatusItem) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	v, err := PathStatusItem(string(bytes))
	if err != nil {
		return err
	}
	*item = v
	return nil
}
func (item *StatusItem) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}
func (item *StatusItem) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, item.String())), nil
}
func (item *StatusItem) UnmarshalJSON(b []byte) error {
	str := strings.ReplaceAll(string(b), "\"", "")
	v, err := PathStatusItem(str)
	if err != nil {
		return common.ErrFoundItem
	}
	*item = v
	return nil
}
