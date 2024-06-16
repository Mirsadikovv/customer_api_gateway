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

	r.GET("/category/getall", handler.GetAllCategory)
	r.GET("/category/get/:id", handler.GetCategoryById)
	r.POST("/category/create", handler.CreateCategory)
	r.PUT("/category/update", handler.UpdateCategory)
	r.DELETE("/category/delete/:id", handler.DeleteCategory)

	r.GET("/product/getall", handler.GetAllProduct)
	r.GET("/product/get/:id", handler.GetProductById)
	r.POST("/product/create", handler.CreateProduct)
	r.PUT("/product/update", handler.UpdateProduct)
	r.DELETE("/product/delete/:id", handler.DeleteProduct)

	r.GET("/review/getall/:id", handler.GetAllReview)
	r.POST("/review/create", handler.CreateReview)
	r.DELETE("/review/delete/:id", handler.DeleteReview)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
