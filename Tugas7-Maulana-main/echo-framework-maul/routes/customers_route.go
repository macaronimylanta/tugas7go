package routes

import (
	"echo-framework-maul/controllers"

	"github.com/labstack/echo"
)

//CustomersRoute ...
func CustomersRoute(g *echo.Group) {

	g.GET("/list", controllers.FetchAllCustomers)

	g.POST("/add", controllers.StoreCustomer)

	g.POST("/update", controllers.UpdateCustomer)

	g.POST("/delete", controllers.DeleteCustomer)

}
