package restaurantstorages

import (
	"context"
	"delivery/modules/restaurants/restaurantmodels"
)

func (s *sqlStore) DeleteItem(
	ctx context.Context,
	id int,
) error {
	db := s.db
	err := db.Table(restaurantmodels.Restaurant{}.TableName()).Where("id = ?", id).Update("status", 0).Error
	if err != nil {
		return err
	}
	return nil
}
