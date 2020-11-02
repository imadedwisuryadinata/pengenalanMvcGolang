package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imadedwisuryadinata/pengenalanMvcGolang/app/controller"
)

func main() {
	router := gin.Default()
	//router.GET("/". getSomething)
	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian", controller.PageAntrianHandler)
	router.Run(":3000")
}

// func getSomething(c *gin.Context){
// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"body1": "Something Success"
// 	})

// return
// }
