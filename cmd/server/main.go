package server

import (
	"armada-api/databases"
	"fmt"

	"github.com/gin-gonic/gin"
)


func Run() {
	db := databases.InitDB()
	defer db.Close()
	
	fmt.Println("Successfully connected!")

	router := gin.Default()

	// router.POST("/login", controllers.Login)
	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// authorized := router.Group("/")
	// authorized.Use(middleware.JWTAuthMiddleware())
	// {
	// 	authorized.GET("/transaction", controllers.GetAllTransactions)
	// 	authorized.POST("/transaction", controllers.InsertTransaction)
	// 	authorized.PUT("/transaction/:id", controllers.UpdateTransaction)
	// 	authorized.DELETE("/transaction/:id", controllers.DeleteTransaction)

	// 	authorized.GET("/category", controllers.GetAllCategories)
	// 	authorized.POST("/category", controllers.InsertCategory)
	// 	authorized.PUT("/category/:id", controllers.UpdateCategory)
	// 	authorized.DELETE("/category/:id", controllers.DeleteCategory)

	// 	authorized.GET("/user", controllers.GetAllUsers)
	// 	authorized.POST("/user", controllers.InsertUser)
	// 	authorized.PUT("/user/:id", controllers.UpdateUser)
	// 	authorized.DELETE("/user/:id", controllers.DeleteUser)

	// 	authorized.GET("/budget", controllers.GetAllBudgets)
	// 	authorized.POST("/budget", controllers.InsertBudget)
	// 	authorized.PUT("/budget/:id", controllers.UpdateBudget)
	// 	authorized.DELETE("/budget/:id", controllers.DeleteBudget)

	// 	router.GET("/budget-status", controllers.GetBudgetStatus)

	// }

	router.Run(":8080")
}