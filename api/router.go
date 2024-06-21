package api

import (
	_ "customer-api-gateway/api/docs" //for swagger
	"customer-api-gateway/api/handler"
	"customer-api-gateway/config"
	"customer-api-gateway/pkg/grpc_client"
	"customer-api-gateway/pkg/logger"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// New ...
// @title           Swagger Category Service API
// @version         1.0
// @description     This is a Catalog service server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	// r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")
	// config.AllowOrigins = cnf.Cfg.AllowOrigins
	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})

	///////////////////////// Catalog_service
	r.GET("/v1/category/getall", handler.GetAllCategory)
	r.GET("/v1/category/get/:id", handler.GetCategoryById)
	r.POST("/v1/category/create", handler.CreateCategory)
	r.PUT("/v1/category/update", handler.UpdateCategory)
	r.DELETE("/v1/category/delete/:id", handler.DeleteCategory)

	r.GET("/v1/product/getall", handler.GetAllProduct)
	r.GET("/v1/product/get/:id", handler.GetProductById)
	r.POST("/v1/product/create", handler.CreateProduct)
	r.PUT("/v1/product/update", handler.UpdateProduct)
	r.DELETE("/v1/product/delete/:id", handler.DeleteProduct)

	r.GET("/v1/review/getall/:id", handler.GetAllReview)
	r.POST("/v1/review/create", handler.CreateReview)
	r.DELETE("/v1/review/delete/:id", handler.DeleteReview)

	/////////////////////// Order_service
	r.POST("/v1/order", handler.CreateOrderProducts)
	r.GET("/v1/order/:id", handler.OrderProductetById)
	r.POST("/v1/orders", handler.CreateOrders)
	r.GET("/v1/orders/:id", handler.OrderstById)

	r.POST("/v1/order-status", handler.CreateOrderstatus)
	r.GET("/v1/order-status/:id", handler.OrderStatusById)
	r.PATCH("/v1/order-status", handler.OrderStatusPatch)
	r.GET("/v1/order-statusp/:id", handler.OrderStatusputch)

	//////////////////////// User_service
	r.POST("/v1/customer/create", handler.CreateCustomer)
	r.GET("/v1/customer/getlist", handler.GetListCustomer)
	r.GET("/v1/customer/getbyid/:id", handler.GetCustomerByID)
	r.PUT("/v1/customer/update/:id", handler.UpdateCustomer)
	r.DELETE("/v1/customer/delete/:id", handler.DeleteCustomer)

	r.POST("/v1/user/create", handler.CreateUser)
	r.GET("/v1/user/getlist", handler.GetListUser)
	r.GET("/v1/user/getbyid/:id", handler.GetUserByID)
	r.PUT("/v1/user/update/:id", handler.UpdateUser)
	r.DELETE("/v1/user/delete/:id", handler.DeleteUser)

	r.POST("/v1/seller/create", handler.CreateSeller)
	r.GET("/v1/seller/getlist", handler.GetListSeller)
	r.GET("/v1/seller/getbyid/:id", handler.GetSellerByID)
	r.PUT("/v1/seller/update/:id", handler.UpdateSeller)
	r.DELETE("/v1/seller/delete/:id", handler.DeleteSeller)

	r.POST("/v1/branch/create", handler.CreateBranch)
	r.GET("/v1/branch/getlist", handler.GetListBranch)
	r.GET("/v1/branch/getbyid/:id", handler.GetBranchByID)
	r.PUT("/v1/branch/update/:id", handler.UpdateBranch)
	r.DELETE("/v1/branch/delete/:id", handler.DeleteBranch)

	r.POST("/v1/shop/create", handler.CreateShop)
	r.GET("/v1/shop/getlist", handler.GetListShop)
	r.GET("/v1/shop/getbyid/:id", handler.GetShopByID)
	r.PUT("/v1/shop/update/:id", handler.UpdateShop)
	r.DELETE("/v1/shop/delete/:id", handler.DeleteShop)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}

// func authMiddleware(c *gin.Context) {
// 	auth := c.GetHeader("Authorization")
// 	if auth == "" {
// 		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
// 	}
// 	c.Next()
// }
