package main

import (
	"github.com/Newbwen/automation-ops/backend/database"
	"github.com/Newbwen/automation-ops/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	i := database.InitProject{}
	i.Init()
	r := gin.Default()
	routes.SetupRoutes()
	r.Run(":8080")
}
