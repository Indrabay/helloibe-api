package middleware

import (
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/utils/logger"
	"go.uber.org/zap"
)

func ValidateStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		// super admin is ignore when validated
		role, _ := c.Get("role")
		if role.(int) == 1 {
			c.Next()
			return
		}
		// in url params, store id is always 1 value for spesific store
		store := c.Param("store_id")
		storeID, err := strconv.ParseInt(store, 10, 64)
		if err != nil {
			logger.Error("error parseInt store_id params",
				zap.String("store", store),
			)
			Unauthorized(c)
			return
		}
		storeVal, _ := c.Get("stores")
		storeArr := storeVal.([]int64)

		if !slices.Contains(storeArr, storeID) {
			logger.Error("user not authorized in the selected store",
				zap.Int64("store_id", storeID),
				zap.Any("storeArr", storeArr),
			)
			Unauthorized(c)
			return
		}

		c.Set("store_id", storeID)

		c.Next()
	}
}
