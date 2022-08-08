package ginrestaurant

import (
	"delivery/common/dataresponse"
	"delivery/component/appctx"
	"delivery/modules/restaurants/restaurantbiz"
	"delivery/modules/restaurants/restaurantstorages"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetDetailRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorages.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetDetailRestaurantBiz(store)
		result, err := biz.GetDetailRestaurant(context.Request.Context(), id)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if result.Id == 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Item has been deleted",
			})
			return
		}
		context.JSON(http.StatusOK, dataresponse.SimpleSuccessResponse(result))
	}
}
