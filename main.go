package main

import (
	"log"

	"github.com/jeromelesaux/fizzbuzz/docs"

	"github.com/gin-gonic/gin"
	"github.com/jeromelesaux/fizzbuzz/configuration"

	"github.com/jeromelesaux/fizzbuzz/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var GitCommit string
var VersionDate string

// @title 					swagger for the fizzbuzz implementation
// @version 			1.0
// @description		A simple fizzbuzz rest api implementation in Golang
// @termsOfService  http://swagger.io/terms/
// @contact.name	Jerome
// @contact.url 	https://github.com/jeromelesaux
// @contact.email	jeromelesaux@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	if err := configuration.CheckConfiguration(); err != nil {
		panic(err.Error())
	}

	r := gin.Default()
	api := r.Group("/api/v1")
	fizzbuzzHandler := handler.NewFizzbuzz()
	api.GET("/fizzbuzz", fizzbuzzHandler.Fizzbuzz)
	api.GET("/stats", fizzbuzzHandler.GetStats)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Default().Printf("Fizzbuzz version [hask: %s, date: %s]\n", GitCommit, VersionDate)

	if err := r.Run(":" + configuration.StaticConfiguration.Port); err != nil {
		panic(err.Error())
	}
}
