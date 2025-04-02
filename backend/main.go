package main

import (
	"github.com/Newbwen/automation-ops/backend/database"
	"github.com/Newbwen/automation-ops/backend/routes"
)

func main() {
	i := database.InitProject{}
	i.Init()
	routes.SetupRoutes().Run("0.0.0.0:8080")
}
