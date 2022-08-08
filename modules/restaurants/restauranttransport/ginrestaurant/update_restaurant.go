package ginrestaurant

import (
	"delivery/common/dataresponse"
	"delivery/component/appctx"
	"delivery/modules/restaurants/restaurantbiz"
	"delivery/modules/restaurants/restaurantmodels"
	"delivery/modules/restaurants/restaurantstorages"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data restaurantmodels.RestaurantUpdate
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorages.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		if err := biz.UpdateRestaurant(context.Request.Context(), id, &data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, dataresponse.SimpleSuccessResponse(true))
	}
}
