package ginrestaurant

import (
	"delivery/common"
	"delivery/common/dataresponse"
	"delivery/component/appctx"
	"delivery/modules/restaurants/restaurantbiz"
	"delivery/modules/restaurants/restaurantmodels"
	"delivery/modules/restaurants/restaurantstorages"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var filter restaurantmodels.Filter
		if err := context.ShouldBind(&filter); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		var paging common.Paging
		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Fulfill()

		store := restaurantstorages.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)

		results, err := biz.ListRestaurants(context.Request.Context(), &filter, &paging)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Fulfill()
		context.JSON(http.StatusOK, dataresponse.NewSuccessResponse(results, paging, filter))
	}
}
