package restaurantbiz

import (
	"context"
	"delivery/modules/restaurants/restaurantmodels"
	"errors"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodels.Restaurant, error)

	DeleteItem(
		ctx context.Context,
		id int,
	) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	item, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil
	}
	if item.Id == 0 {
		return errors.New("Item not found")
	}
	if err := biz.store.DeleteItem(ctx, id); err != nil {
		return err
	}
	return nil
}
