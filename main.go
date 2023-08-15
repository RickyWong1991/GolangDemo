package main

import (
	"net/http"
	"tutorial/golangdemo/controller"
	docs "tutorial/golangdemo/docs"
	"tutorial/golangdemo/model"
	"tutorial/golangdemo/model/database"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@BasePath	/api/v1

// PingExample godoc
//
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/example/helloworld [get]
func HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World")
}

func LoadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Account{})
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	r := gin.Default()
	// model.ConnectDatabase()
	LoadDatabase()
	c := controller.NewController()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", HelloWorld)
		}

		accounts := v1.Group("/accounts")
		{
			// accounts.GET(":id", c.ShowAccount)
			accounts.GET(":name", c.FindUserByUserName)
			accounts.GET("", c.ListAccounts)
			accounts.POST("", c.AddAccount)
			accounts.DELETE(":id", c.DeleteAccount)
			accounts.PATCH(":id", c.UpdateAccount)
			accounts.POST(":id/images", c.UploadAccountImage)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
