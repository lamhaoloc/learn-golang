package gin_restaurant

import (
	"awesomeProject/component/appctx"
	restaurantModel "awesomeProject/module/restaurant/model"
	restaurantService "awesomeProject/module/restaurant/service"
	restaurantStorage "awesomeProject/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data restaurantModel.RestaurantCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantStorage.NewSQLStore(db)
		service := restaurantService.NewCreateRestaurantService(store)

		if err := service.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
