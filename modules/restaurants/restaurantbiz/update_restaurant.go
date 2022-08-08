package restaurantbiz

import (
	"context"
	"delivery/modules/restaurants/restaurantmodels"
	"errors"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodels.Restaurant, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *restaurantmodels.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodels.RestaurantUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Id == 0 {
		return errors.New("Item not found")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
