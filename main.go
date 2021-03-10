package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/sum", handleSum)

	port := os.Getenv("PORT")
	router.Run(":"+port)
}

func handleSum(ctx *gin.Context) {

	// res := doSum(5, 6)

	ctx.JSON(http.StatusOK, "I dey here!")

}

func doSum(a, b int) int {
	return a + b
}