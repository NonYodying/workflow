package model

import "github.com/NonYodying/workflow/internal/constant"

type RequestCreateItem struct {
	Title    string `json:"title" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type RequestUpdateItem struct {
	Title    *string `json:"title"`
	Amount   *int    `json:"amount"`
	Quantity *int    `json:"quantity"`
}
type RequestFindItem struct {
	Statuses constant.ItemStatus `form:"status"`
}

type RequestUpdateStatus struct {
	Status constant.ItemStatus
}

type RequestPatchItemStatus struct {
	Status constant.ItemStatus `json:"status" binding:"required"`
}

type RequestLogin struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
