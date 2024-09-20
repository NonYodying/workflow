package model

import "github.com/NonYodying/workflow/internal/constant"

/***type Item struct {
	ID       uint                `json:"id" gorm:"primaryKey"`
	Title    string              `json:"title"`
	Amount   uint                `json:"amount"`
	Quantity uint                `json:"quantity"`
	Status   constant.ItemStatus `json:"status"`
}***/

type Item struct {
	ID       uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	Title    string              `gorm:"size:255;not null" json:"title"`
	Amount   int                 `gorm:"not null" json:"amount"`
	Quantity int                 `gorm:"not null" json:"quantity"`
	Status   constant.ItemStatus `gorm:"size:20;not null" json:"status"`
}
