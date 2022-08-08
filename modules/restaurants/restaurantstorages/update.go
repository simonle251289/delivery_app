package restaurantstorages

import (
	"context"
	"delivery/modules/restaurants/restaurantmodels"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *restaurantmodels.RestaurantUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
