package restaurantstorages

import (
	"context"
	"delivery/common"
	"delivery/modules/restaurants/restaurantmodels"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodels.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodels.Restaurant, error) {
	var results []restaurantmodels.Restaurant
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	db = db.Table(restaurantmodels.Restaurant{}.TableName()).Where(conditions)
	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
	}
	if err := db.Table(restaurantmodels.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
