package item

import (
	"github.com/NonYodying/workflow/internal/model"

	"github.com/NonYodying/workflow/internal/constant"
	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
}

func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db),
	}
}

func (service Service) Create(req model.RequestCreateItem) (model.Item, error) {
	// Find user id that make request to fill in owner_id

	// Create item
	item := model.Item{
		Title:    req.Title,
		Amount:   req.Amount,
		Quantity: req.Quantity,
		Status:   constant.ItemPendingStatus,
	}

	if err := service.Repository.Create(&item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (service Service) Find(query model.RequestFindItem) ([]model.Item, error) {
	return service.Repository.Find(query)
}

func (service Service) FindByID(id uint) (model.Item, error) {
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (service Service) FindLatestID(query model.RequestFindItem) (model.Item, error) {

	return service.Repository.FindLatestID(query)
}

func (service Service) UpdateLatestItem(query model.RequestFindItem, req model.RequestUpdateItem) (model.Item, error) {

	item, err := service.Repository.FindLatestID(query)
	if err != nil {
		return model.Item{}, err
	}

	if req.Title != nil {
		item.Title = *req.Title
	}
	if req.Amount != nil {
		item.Amount = *req.Amount
	}
	if req.Quantity != nil {
		item.Quantity = *req.Quantity
	}

	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}

	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (service Service) UpdateLatestStatus(query model.RequestFindItem, status constant.ItemStatus) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindLatestID(query)
	if err != nil {
		return model.Item{}, err
	}

	// Fill data
	item.Status = status

	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (service Service) DeleteLatestItem(query model.RequestFindItem) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindLatestID(query)
	if err != nil {
		return model.Item{}, err
	}

	// Delete
	if err := service.Repository.DeleteLatest(item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (service Service) UpdateItem(id uint, req model.RequestUpdateItem) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err
	}

	// Fill data
	if req.Title != nil {
		item.Title = *req.Title
	}
	if req.Amount != nil {
		item.Amount = *req.Amount
	}
	if req.Quantity != nil {
		item.Quantity = *req.Quantity
	}

	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (service Service) UpdateStatus(id uint, status constant.ItemStatus) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err
	}

	// Fill data
	item.Status = status

	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (service Service) Delete(id uint) error {
	return service.Repository.Delete(id)
}
