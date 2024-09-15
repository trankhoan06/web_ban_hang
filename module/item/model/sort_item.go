package model

type CreateSortItem struct {
	Id       int    `json:"id" gorm:"column:id"`
	ItemId   int    `json:"item_id" gorm:"column:item_id"`
	SortItem string `json:"sort_item" gorm:"sort_item"`
	Size     string `json:"size" gorm:"column:size"`
	Price    int    `json:"price" gorm:"column:price"`
	Amount   int    `json:"amount" gorm:"column:amount"`
	Image    *Image `json:"image" gorm:"column:image"`
}
type UpdateSortItem struct {
	Id       int     `json:"id" gorm:"column:id"`
	ItemId   int     `json:"-" gorm:"column:item_id"`
	SortItem *string `json:"sort_item" gorm:"sort_item"`
	Size     *string `json:"size" gorm:"column:size"`
	Price    *int    `json:"price" gorm:"column:price"`
	Amount   *int    `json:"amount" gorm:"column:amount"`
	Image    *Image  `json:"image" gorm:"column:image"`
}
type SortItem struct {
	Id       int    `json:"id" gorm:"column:id"`
	ItemId   int    `json:"item_id" gorm:"column:item_id"`
	SortItem string `json:"sort_item" gorm:"sort_item"`
	Size     string `json:"size" gorm:"column:size"`
	Price    int    `json:"price" gorm:"column:price"`
	Amount   int    `json:"amount" gorm:"column:amount"`
}

func (item *CreateSortItem) TableName() string { return "sort_item" }
func (item *UpdateSortItem) TableName() string { return "sort_item" }
