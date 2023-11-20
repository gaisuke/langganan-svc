package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// db := mysql.DB()

	r.Run()
}
