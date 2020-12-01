package external

import (
	"aggregation-mod/pkg/adapter/controllers"
	"aggregation-mod/pkg/external/pg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	logger := &Logger{}
	conn := pg.Connect()
	experimentController := controllers.NewExperimentController(conn, logger)
	resultController := controllers.NewResultController(conn, logger)

	router.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "hello world") })
	router.GET("/experiments", func(c *gin.Context) { experimentController.Index(c) })
	router.POST("/experiments", func(c *gin.Context) { experimentController.Create(c) })
	router.GET("/experiments/:id", func(c *gin.Context) { experimentController.Show(c) })
	router.DELETE("/experiments/:id", func(c *gin.Context) { experimentController.Delete(c) })

	router.GET("/results", func(c *gin.Context) { resultController.Index(c) })
	router.GET("/results/:id", func(c *gin.Context) { resultController.Show(c) })
	router.POST("/experiments/:experiment_id/results", func(c *gin.Context) { resultController.Create(c) })
	router.DELETE("/results/:id", func(c *gin.Context) { resultController.Delete(c) })

	Router = router
}
