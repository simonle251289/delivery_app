package restaurantbiz

import (
	"context"
	"delivery/modules/restaurants/restaurantmodels"
)

type GetDetailRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodels.Restaurant, error)
}

type getDetailRestaurantBiz struct {
	store GetDetailRestaurantStore
}

func NewGetDetailRestaurantBiz(store GetDetailRestaurantStore) *getDetailRestaurantBiz {
	return &getDetailRestaurantBiz{
		store: store,
	}
}

func (biz getDetailRestaurantBiz) GetDetailRestaurant(ctx context.Context, id int) (*restaurantmodels.Restaurant, error) {

	return biz.store.FindDataByCondition(ctx, map[string]interface{}{"status": 1, "id": id})
}
