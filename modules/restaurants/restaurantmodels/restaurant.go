package restaurantmodels

import (
	"delivery/common"
	"errors"
	"strings"
)

type Restaurant struct {
	common.BaseModel `json:",inline"` //Dùng inline để trả về value cùng cấp với các field hiện tại
	Name             string           `json:"name" gorm:"column:name;"`
	Addr             string           `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (biz *RestaurantCreate) Validate() error {
	biz.Name = strings.TrimSpace(biz.Name)

	if len(biz.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}
	return nil
}
