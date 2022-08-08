package restaurantbiz

import (
	"context"
	"delivery/common"
	"delivery/modules/restaurants/restaurantmodels"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodels.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodels.Restaurant, error)
}
type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz listRestaurantBiz) ListRestaurants(
	ctx context.Context,
	filter *restaurantmodels.Filter,
	paging *common.Paging,
) ([]restaurantmodels.Restaurant, error) {

	results, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	return results, err
}
