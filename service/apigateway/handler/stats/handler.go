package statshandler

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/apierror"
	gateway_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache/gateway_cache"
	pbcategory "github.com/MamangRust/microservice-ecommerce-grpc-pb/category"
	pborder "github.com/MamangRust/microservice-ecommerce-grpc-pb/order"
	pbtransaction "github.com/MamangRust/microservice-ecommerce-grpc-pb/transaction"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	statsmapper "github.com/MamangRust/microservice-ecommerce-shared/mapper/stats"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type DepsStats struct {
	Client     *grpc.ClientConn
	Router     chi.Router
	Logger     logger.LoggerInterface
	CacheStore *cache.CacheStore
	ApiHandler apierror.ApiHandler
}

// RegisterStatsHandler exposes every stats-reader (ClickHouse) RPC as a REST
// endpoint. The stats-reader registers its services on a single gRPC server, so
// all clients are built from the same connection.
func RegisterStatsHandler(deps *DepsStats) {
	mapper := statsmapper.NewStatsResponseMapper()
	cache := gateway_cache.NewGatewayCache(deps.CacheStore)

	NewStatsHandleApi(&statsHandleDeps{
		orderClient:           pborder.NewOrderStatsServiceClient(deps.Client),
		orderByMerchantClient: pborder.NewOrderStatsByMerchantServiceClient(deps.Client),
		categoryClient:        pbcategory.NewCategoryStatsServiceClient(deps.Client),
		categoryByIdClient:    pbcategory.NewCategoryStatsByIdServiceClient(deps.Client),
		categoryByMerchantCli: pbcategory.NewCategoryStatsByMerchantServiceClient(deps.Client),
		transactionClient:     pbtransaction.NewTransactionStatsServiceClient(deps.Client),
		transactionByMerchant: pbtransaction.NewTransactionStatsByMerchantServiceClient(deps.Client),
		router:                deps.Router,
		logger:                deps.Logger,
		mapping:               mapper,
		cache:                 cache,
		apiHandler:            deps.ApiHandler,
	})
}
