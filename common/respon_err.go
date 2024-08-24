package common

import (
	"errors"
)

var (
	ErrFoundItem = errors.New("you don't have this item")
	ErrFindItem  = errors.New("item has been deleted of no exist")
	ErrNoPermiss = errors.New("no permiss")
)
