package router

import (
	"github.com/ArtyDev57/lxd-warehouse/api/handler"
	"github.com/labstack/echo"
)

// GetAllRoutes function
func GetAllRoutes(e *echo.Echo) {

	product := e.Group("/api/product")

	product.POST("/create", handler.CreateProduct)

	product.GET("/all", handler.GetAllProducts)

	product.GET("/:id", handler.GetProduct)

	product.PUT("/edit/:id", handler.UpdateProduct)

	product.GET("/import/all", handler.GetAllImport)

	product.POST("/import/order", handler.ImportProduct)

	product.POST("/import/order/detail/:id", handler.ImportProductDetail)

	product.GET("/import/detail/:id", handler.GetImportDetail)

	product.PUT("/import/confirm/:id", handler.ConfirmImport)

	product.GET("/export/all", handler.GetAllExport)

	product.POST("/export", handler.ExportProduct)

	product.POST("/export/detail/:id", handler.ExportProductDetail)

	product.GET("/export/detail/:id", handler.GetExportDetail)

	product.PUT("/export/confirm/:id", handler.ConfirmExport)

	extras := e.Group("/api/extras")

	extras.POST("/product-brand/create", handler.CreateProductBrand)

	extras.PUT("/product-brand/edit/:id", handler.UpdateProductBrand)

	extras.GET("/product-brand/all", handler.GetProductBrands)

	extras.POST("/product-type/create", handler.CreateProductType)

	extras.PUT("/product-type/edit/:id", handler.UpdateProductType)

	extras.GET("/product-type/all", handler.GetProductTypes)

	extras.POST("/supplier/create", handler.CreateSupplier)

	extras.PUT("/supplier/edit/:id", handler.UpdateSupplier)

	extras.GET("/supplier/all", handler.GetSuppliers)

	extras.POST("/distributor/create", handler.CreateDistributor)

	extras.PUT("/distributor/edit/:id", handler.UpdateDistributor)

	extras.GET("/distributor/all", handler.GetDistributors)

	users := e.Group("/api/user")

	users.POST("/create", handler.CreateUser)

	users.POST("/login", handler.Login)

	users.GET("/all", handler.GetAllUser)

	users.GET("/:username", handler.GetUser)

	users.GET("/role/all", handler.GetAllUserRole)

	users.PUT("/update/:id", handler.UpdateUser)
}
