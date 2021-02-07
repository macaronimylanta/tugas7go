package routes

import (
	"echo-framework-maul/controllers"

	"github.com/labstack/echo"
)

func SuppliersRoute(g *echo.Group) {
	g.GET("/list", controllers.FetchAllSuppliers)

	g.POST("/add", controllers.StoreSupplier)

	g.PUT("/update/:supplierID", controllers.UpdateSupplier)

	g.DELETE("/delete/:supplierID", controllers.DeleteSupplier)
}
