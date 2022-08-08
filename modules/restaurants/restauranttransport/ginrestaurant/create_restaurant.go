package ginrestaurant

import (
	"delivery/common/dataresponse"
	"delivery/component/appctx"
	"delivery/modules/restaurants/restaurantbiz"
	"delivery/modules/restaurants/restaurantmodels"
	"delivery/modules/restaurants/restaurantstorages"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data restaurantmodels.RestaurantCreate
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorages.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		err := biz.CreateRestaurant(context.Request.Context(), &data)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, dataresponse.SimpleSuccessResponse(data))
	}
}
