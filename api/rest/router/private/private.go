package private

import (
	"github.com/HiBang15/golang-api-base-tulpo.git"
	"github.com/gin-gonic/gin"
	"log"
)

func SetRouter(router *gin.RouterGroup) {
	log.Print("Start init private router .....")
	// Define all private router in here
	//middleware auth
	//authMiddleware, err := auth.New(constant.SIGNING_ALGORITHM, constant.SECRET_KEY, constant.TOKEN_TIMEOUT)
	//if err != nil {
	//	log.Fatalf("has error when init auth middleware: %v", err)
	//}
	////group Auth
	//auth := router.Group("auth")
	//auth.Use(authMiddleware.MiddlewareFunc())
	//auth.GET("me", controller.AuthMe)
	//auth.POST("update")
	//auth.POST("refresh-token", controller.RefreshToken)
	//auth.POST("logout", controller.Logout)
	//auth.POST("logout-all", controller.LogoutAllDevice)


	//group User
	user := router.Group("user")
	{
		user.POST("/create", controllers.CreateUserAccount)
		//user.POST("/", controllers.CreateUserAccount)
		////user.GET("get-list-account")
		//user.PUT("/:id", controllers.UpdateUserAccount)
		//user.DELETE("/:id", controllers.DeleteUserAccount)
		//user.GET("/all", controllers.GetAllUser)
	}

	log.Print("Finish init private router ....")
}