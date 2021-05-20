package router

import (
	"fmt"
	"github.com/HiBang15/golang-api-base-tulpo.git/api/rest/router/private"
	"github.com/HiBang15/golang-api-base-tulpo.git/api/rest/router/public"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

var (
	ListenAddress string
)

func init() {
	ListenAddress = fmt.Sprintf(":%s", os.Getenv("LISTEN_ADDRESS_PORT"))
	if os.Getenv("LISTEN_ADDRESS_PORT") == "" {
		ListenAddress = "0.0.0.0:8080"
	}
}

func Start(environment string) {
	// run mode
	switch environment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
		fmt.Printf("Start product mode...\nServer Listen on: %v", ListenAddress)
		fmt.Println()
	}
	router := gin.New()

	// setting router
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:           []string{"*"},
		AllowMethods:           []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders:           []string{"Origin", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
		AllowCredentials:       false,
		ExposeHeaders:          []string{"Content-Length"},
		MaxAge:                 12 * time.Hour,
	}))

	basePath := os.Getenv("API_VERSION")
	apiRouters := router.Group(basePath)
	//set public router
	public.SetRouter(apiRouters)
	//set private router
	private.SetRouter(apiRouters)

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//run
	err := router.Run(ListenAddress)
	log.Println(err)
}