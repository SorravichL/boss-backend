package main

import (
    "boss-backend/internal/config"
    "boss-backend/internal/db"
    "boss-backend/internal/handlers"
    "boss-backend/internal/models"

    "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
    // auto-migrate (dev only)
    db.AutoMigrate(&models.Customer{}, &models.Supplier{}, &models.Employee{})

    r := gin.Default()
    custH := handlers.NewCustomerHandler(db)
	supH := handlers.NewSupplierHandler(db)
	empH := handlers.NewEmployeeHandler(db)

    // Group all /customers routes
    c := r.Group("/customers")
    {
        c.GET("", custH.List)
        c.POST("", custH.Create)
        c.GET("/:id", custH.Get)
        c.PUT("/:id", custH.Update)
        c.DELETE("/:id", custH.Delete)
    }

	sGroup := r.Group("/suppliers")
    {
        sGroup.GET("", supH.List)
        sGroup.POST("", supH.Create)
        sGroup.GET(":id", supH.Get)
        sGroup.PUT(":id", supH.Update)
        sGroup.DELETE(":id", supH.Delete)
    }
	eGroup := r.Group("/employees")
	{
		eGroup.GET("", empH.List)
		eGroup.POST("", empH.Create)
		eGroup.GET(":id", empH.Get)
		eGroup.PUT(":id", empH.Update)
		eGroup.DELETE(":id", empH.Delete)
	}

    return r
}

func main() {
    cfg := config.Load()
    gormDB := db.NewGormDB(cfg)

    router := setupRouter(gormDB)
    router.Run(":8080")
}
