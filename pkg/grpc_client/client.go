package grpc_client

import (
	"customer-api-gateway/config"
	pc "customer-api-gateway/genproto/catalog_service"
	pd "customer-api-gateway/genproto/product_service"
	rv "customer-api-gateway/genproto/review_service"
	"log"

	op "customer-api-gateway/genproto/order_product"
	os "customer-api-gateway/genproto/order_status_notes"
	ct "customer-api-gateway/genproto/orders_service"

	us "customer-api-gateway/genproto/user_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	CategoryService() pc.CategoryServiceClient
	ProductService() pc.CategoryServiceClient
	ReviewService() rv.ReviewServiceClient

	ProducOrderService() op.OrderProductsServiceClient
	OrderService() ct.OrderServiceServer
	OrderStatus() os.OrderStatusServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {
	connCatalog, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port:%s err: %s",
			cfg.CatalogServiceHost, cfg.CatalogServicePort, err)
	}
	connOrder, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("order service dial hsot:%s port :%s err:%s",
			cfg.OrderServiceHost, cfg.OrderServicePort, err)
	}
	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"category_service":     pc.NewCategoryServiceClient(connCatalog),
			"product_service":      pd.NewProductServiceClient(connCatalog),
			"review_service":       rv.NewReviewServiceClient(connCatalog),
			"orderproduct_service": op.NewOrderProductsServiceClient(connOrder),
			"order_service":        ct.NewOrderServiceClient(connOrder),
			"order_status":         os.NewOrderStatusServiceClient(connOrder),
		},
	}, nil
}

func (g *GrpcClient) CategoryService() pc.CategoryServiceClient {
	return g.connections["category_service"].(pc.CategoryServiceClient)
}

func (g *GrpcClient) ProductService() pd.ProductServiceClient {
	return g.connections["product_service"].(pd.ProductServiceClient)
}

func (g *GrpcClient) ReviewService() rv.ReviewServiceClient {
	return g.connections["review_service"].(rv.ReviewServiceClient)
}

func (o *GrpcClient) ProducOrderService() op.OrderProductsServiceClient {
	return o.connections["orderproduct_service"].(op.OrderProductsServiceClient)
}

func (o *GrpcClient) OrderService() ct.OrderServiceClient {
	return o.connections["order_service"].(ct.OrderServiceClient)
}

func (o *GrpcClient) OrderStatus() os.OrderStatusServiceClient {
	return o.connections["order_status"].(os.OrderStatusServiceClient)
}

func (g *GrpcClient) UserService() us.CustomerServiceClient {
	client, ok := g.connections["user_service"].(us.CustomerServiceClient)
	if !ok {
		log.Println("failed to assert type for user_service")
		return nil
	}
	return client
}

func (g *GrpcClient) SystemUserService() us.UsServiceClient {
	client, ok := g.connections["system_user"].(us.UsServiceClient)
	if !ok {
		log.Println("failed to assert type for system_user")
		return nil
	}
	return client
}

func (g *GrpcClient) SellerService() us.SellerServiceClient {
	client, ok := g.connections["seller"].(us.SellerServiceClient)
	if !ok {
		log.Println("failed to assert type for seller")
		return nil
	}
	return client
}

func (g *GrpcClient) BranchService() us.BranchServiceClient {
	client, ok := g.connections["branch"].(us.BranchServiceClient)
	if !ok {
		log.Println("failed to assert type for branch")
		return nil
	}
	return client
}

func (g *GrpcClient) ShopService() us.ShopServiceClient {
	client, ok := g.connections["shop"].(us.ShopServiceClient)
	if !ok {
		log.Println("failed to assert type for shop")
		return nil
	}
	return client
}
