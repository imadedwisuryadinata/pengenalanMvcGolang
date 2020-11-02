package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.GET("/". getSomething)
	router.POST("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/status", GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrian/delete", DeleteAntrianHandler)
	router.Run(":3000")
}

// func getSomething(c *gin.Context){
// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"body1": "Something Success"
// 	})

// return
// }
