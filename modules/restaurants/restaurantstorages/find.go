package restaurantstorages

import (
	"context"
	"delivery/modules/restaurants/restaurantmodels"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodels.Restaurant, error) {
	var result restaurantmodels.Restaurant
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	db = db.Table(restaurantmodels.Restaurant{}.TableName()).Where(conditions)
	if err := db.
		Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
