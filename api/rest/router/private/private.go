package private

import (
	"log"

	"github.com/HiBang15/golang-api-base-tulpo.git/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.RouterGroup) {
	log.Print("Start init private router .....")

	//group User
	user := router.Group("user")
	{
		user.POST("/create", controllers.CreateUserAccount)
	}

	//group activity
	activity := router.Group("activity")
	{
		activity.POST("/create", controllers.CreateActivity)
		activity.PUT("/update", controllers.UpdateActivity)
		activity.DELETE("/delete/:id", controllers.DeleteActivity)
		activity.GET("/search/:id", controllers.GetActivityByID)
	}

	//permission
	permission := router.Group("/permission")
	{
		permission.POST("/create", controllers.CreatePermission)
	}

	log.Print("Finish init private router ....")

	// permission := router.Group("permission")
}
