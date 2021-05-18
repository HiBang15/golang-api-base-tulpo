package private

import (
	"github.com/HiBang15/golang-api-base-tulpo.git/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func SetRouter(router *gin.RouterGroup) {
	log.Print("Start init private router .....")


	//group User
	user := router.Group("user")
	{
		user.POST("/create", controllers.CreateUserAccount)
	}

	log.Print("Finish init private router ....")
}