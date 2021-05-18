package public

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SetRouter(router *gin.RouterGroup) {
	log.Print("Start init public router .....")
	// Define all public route in here
	//router.GET("/", cache.Cached(12, func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{"message": "Say hello from SmartOSC", "data": "nothing!!!"})
	//}))


	//register := router.Group("register")
	//{
	//	register.POST("/account", controller.CreateUserAccount)
	//}
	//
	////router.POST("/auth/login", controller.Login)
	//
	////group Auth
	//auth := router.Group("auth")
	//auth.POST("/login", controller.Login)
	//auth.POST("forgot-password/send-code", controller.SendCode)
	//auth.POST("forgot-password/confirm-code", controller.ConfirmCode)
	//auth.POST("forgot-password/new-password", controller.NewPassword)
	//
	//verify := router.Group("verify-account")
	//verify.POST("/", controller.VerifyEmail)

	log.Print("Finish init public router ....")
}
