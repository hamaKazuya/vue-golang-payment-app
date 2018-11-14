package infrastructure

import (
	"log"
	"vue-golang-payment-app/backend-api/handler"

	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

// Router ほげほげ
var Router *gin.Engine

func init() {
	router := gin.Default()
	log.SetFlags(log.Llongfile)
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	router.GET("/api/v1/items", func(c *gin.Context) { handler.GetLists(c) })
	router.GET("/api/v1/items/:id", func(c *gin.Context) { handler.GetItem(c) })
	router.POST("/api/v1/charge/items/:id", func(c *gin.Context) { handler.Charge(c) })

	Router = router
}
