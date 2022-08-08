package restaurantstorages

import (
	"context"
	"delivery/modules/restaurants/restaurantmodels"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodels.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
