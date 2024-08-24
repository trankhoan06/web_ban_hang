package model

import (
	"strings"
	"time"
)

type StatusSearch int

const (
	StatusSearchInactive StatusSearch = iota
	StatusSearchActive
)

type Search struct {
	Id         int          `json:"id" gorm:"column:id"`
	Content    string       `json:"content" gorm:"column:content"`
	SearchTime *int         `json:"-" gorm:"column:search_time"`
	Status     StatusSearch `json:"-" gorm:"column:status"`
	CreateAt   time.Time    `json:"-" gorm:"column:create_at"`
	UpdateAt   time.Time    `json:"update_at" gorm:"column:update_at"`
}
type SearchKeyword struct {
	Content    string    `json:"content" gorm:"column:content"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
}
type CategorySearch struct {
	Name        string `json:"name"`
	Arrangement string `json:"arrangement"`
}

func (Category CategorySearch) Process() {
	Category.Arrangement = strings.TrimSpace(Category.Arrangement)
	Category.Name = strings.TrimSpace(Category.Name)
	if Category.Arrangement == "" {
		Category.Arrangement = "DESC"
	}
}

func (Search) TableName() string { return "search" }
