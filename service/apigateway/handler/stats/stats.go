package statshandler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/apierror"
	gateway_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache/gateway_cache"
	pbcategory "github.com/MamangRust/microservice-ecommerce-grpc-pb/category"
	pborder "github.com/MamangRust/microservice-ecommerce-grpc-pb/order"
	pbtransaction "github.com/MamangRust/microservice-ecommerce-grpc-pb/transaction"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	sharedErrors "github.com/MamangRust/microservice-ecommerce-shared/errors"
	statsmapper "github.com/MamangRust/microservice-ecommerce-shared/mapper/stats"
	"github.com/go-chi/chi/v5"
)

const statsCacheTTL = 5 * time.Minute

// statsHandleApi proxies the stats-reader gRPC contracts as REST endpoints.
type statsHandleApi struct {
	orderClient           pborder.OrderStatsServiceClient
	orderByMerchantClient pborder.OrderStatsByMerchantServiceClient

	categoryClient        pbcategory.CategoryStatsServiceClient
	categoryByIdClient    pbcategory.CategoryStatsByIdServiceClient
	categoryByMerchantCli pbcategory.CategoryStatsByMerchantServiceClient

	transactionClient     pbtransaction.TransactionStatsServiceClient
	transactionByMerchant pbtransaction.TransactionStatsByMerchantServiceClient

	logger     logger.LoggerInterface
	mapping    statsmapper.StatsResponseMapper
	cache      *gateway_cache.GatewayCache
	apiHandler apierror.ApiHandler
}

type statsHandleDeps struct {
	orderClient           pborder.OrderStatsServiceClient
	orderByMerchantClient pborder.OrderStatsByMerchantServiceClient

	categoryClient        pbcategory.CategoryStatsServiceClient
	categoryByIdClient    pbcategory.CategoryStatsByIdServiceClient
	categoryByMerchantCli pbcategory.CategoryStatsByMerchantServiceClient

	transactionClient     pbtransaction.TransactionStatsServiceClient
	transactionByMerchant pbtransaction.TransactionStatsByMerchantServiceClient

	router     chi.Router
	logger     logger.LoggerInterface
	mapping    statsmapper.StatsResponseMapper
	cache      *gateway_cache.GatewayCache
	apiHandler apierror.ApiHandler
}

func NewStatsHandleApi(params *statsHandleDeps) *statsHandleApi {
	statsHandler := &statsHandleApi{
		orderClient:           params.orderClient,
		orderByMerchantClient: params.orderByMerchantClient,
		categoryClient:        params.categoryClient,
		categoryByIdClient:    params.categoryByIdClient,
		categoryByMerchantCli: params.categoryByMerchantCli,
		transactionClient:     params.transactionClient,
		transactionByMerchant: params.transactionByMerchant,
		logger:                params.logger,
		mapping:               params.mapping,
		cache:                 params.cache,
		apiHandler:            params.apiHandler,
	}

	// Order stats live under /api/order to match the documented contract
	// (/api/order-command and /api/order-query own the CRUD routes).
	params.router.Route("/api/order", func(orderStats chi.Router) {
		orderStats.Get("/monthly-total-revenue", params.apiHandler.Handle("get-order-stats-monthly-total-revenue", statsHandler.FindOrderMonthlyTotalRevenue))
		orderStats.Get("/yearly-total-revenue", params.apiHandler.Handle("get-order-stats-yearly-total-revenue", statsHandler.FindOrderYearlyTotalRevenue))
		orderStats.Get("/monthly-revenue", params.apiHandler.Handle("get-order-stats-monthly-revenue", statsHandler.FindOrderMonthlyRevenue))
		orderStats.Get("/yearly-revenue", params.apiHandler.Handle("get-order-stats-yearly-revenue", statsHandler.FindOrderYearlyRevenue))

		orderStats.Route("/merchant", func(byMerchant chi.Router) {
			byMerchant.Get("/monthly-total-revenue", params.apiHandler.Handle("get-order-stats-merchant-monthly-total-revenue", statsHandler.FindOrderMonthlyTotalRevenueByMerchant))
			byMerchant.Get("/yearly-total-revenue", params.apiHandler.Handle("get-order-stats-merchant-yearly-total-revenue", statsHandler.FindOrderYearlyTotalRevenueByMerchant))
			byMerchant.Get("/monthly-revenue", params.apiHandler.Handle("get-order-stats-merchant-monthly-revenue", statsHandler.FindOrderMonthlyRevenueByMerchant))
			byMerchant.Get("/yearly-revenue", params.apiHandler.Handle("get-order-stats-merchant-yearly-revenue", statsHandler.FindOrderYearlyRevenueByMerchant))
		})
	})

	params.router.Route("/api/category-stats", func(categoryStats chi.Router) {
		categoryStats.Get("/monthly-total-pricing", params.apiHandler.Handle("get-category-stats-monthly-total-pricing", statsHandler.FindCategoryMonthlyTotalPrice))
		categoryStats.Get("/yearly-total-pricing", params.apiHandler.Handle("get-category-stats-yearly-total-pricing", statsHandler.FindCategoryYearlyTotalPrice))
		categoryStats.Get("/monthly-pricing", params.apiHandler.Handle("get-category-stats-monthly-pricing", statsHandler.FindCategoryMonthPrice))
		categoryStats.Get("/yearly-pricing", params.apiHandler.Handle("get-category-stats-yearly-pricing", statsHandler.FindCategoryYearPrice))

		categoryStats.Route("/mycategory", func(byId chi.Router) {
			byId.Get("/monthly-total-pricing", params.apiHandler.Handle("get-category-stats-mycategory-monthly-total-pricing", statsHandler.FindCategoryMonthlyTotalPriceById))
			byId.Get("/yearly-total-pricing", params.apiHandler.Handle("get-category-stats-mycategory-yearly-total-pricing", statsHandler.FindCategoryYearlyTotalPriceById))
			byId.Get("/monthly-pricing", params.apiHandler.Handle("get-category-stats-mycategory-monthly-pricing", statsHandler.FindCategoryMonthPriceById))
			byId.Get("/yearly-pricing", params.apiHandler.Handle("get-category-stats-mycategory-yearly-pricing", statsHandler.FindCategoryYearPriceById))
		})

		categoryStats.Route("/merchant", func(byMerchant chi.Router) {
			byMerchant.Get("/monthly-total-pricing", params.apiHandler.Handle("get-category-stats-merchant-monthly-total-pricing", statsHandler.FindCategoryMonthlyTotalPriceByMerchant))
			byMerchant.Get("/yearly-total-pricing", params.apiHandler.Handle("get-category-stats-merchant-yearly-total-pricing", statsHandler.FindCategoryYearlyTotalPriceByMerchant))
			byMerchant.Get("/monthly-pricing", params.apiHandler.Handle("get-category-stats-merchant-monthly-pricing", statsHandler.FindCategoryMonthPriceByMerchant))
			byMerchant.Get("/yearly-pricing", params.apiHandler.Handle("get-category-stats-merchant-yearly-pricing", statsHandler.FindCategoryYearPriceByMerchant))
		})
	})

	params.router.Route("/api/transaction-stats", func(transactionStats chi.Router) {
		transactionStats.Get("/monthly-success", params.apiHandler.Handle("get-transaction-stats-monthly-success", statsHandler.FindTransactionMonthlySuccess))
		transactionStats.Get("/yearly-success", params.apiHandler.Handle("get-transaction-stats-yearly-success", statsHandler.FindTransactionYearlySuccess))
		transactionStats.Get("/monthly-failed", params.apiHandler.Handle("get-transaction-stats-monthly-failed", statsHandler.FindTransactionMonthlyFailed))
		transactionStats.Get("/yearly-failed", params.apiHandler.Handle("get-transaction-stats-yearly-failed", statsHandler.FindTransactionYearlyFailed))
		transactionStats.Get("/monthly-method-success", params.apiHandler.Handle("get-transaction-stats-monthly-method-success", statsHandler.FindTransactionMonthlyMethodSuccess))
		transactionStats.Get("/yearly-method-success", params.apiHandler.Handle("get-transaction-stats-yearly-method-success", statsHandler.FindTransactionYearlyMethodSuccess))
		transactionStats.Get("/monthly-method-failed", params.apiHandler.Handle("get-transaction-stats-monthly-method-failed", statsHandler.FindTransactionMonthlyMethodFailed))
		transactionStats.Get("/yearly-method-failed", params.apiHandler.Handle("get-transaction-stats-yearly-method-failed", statsHandler.FindTransactionYearlyMethodFailed))

		transactionStats.Route("/merchant", func(byMerchant chi.Router) {
			byMerchant.Get("/monthly-success", params.apiHandler.Handle("get-transaction-stats-merchant-monthly-success", statsHandler.FindTransactionMonthlySuccessByMerchant))
			byMerchant.Get("/yearly-success", params.apiHandler.Handle("get-transaction-stats-merchant-yearly-success", statsHandler.FindTransactionYearlySuccessByMerchant))
			byMerchant.Get("/monthly-failed", params.apiHandler.Handle("get-transaction-stats-merchant-monthly-failed", statsHandler.FindTransactionMonthlyFailedByMerchant))
			byMerchant.Get("/yearly-failed", params.apiHandler.Handle("get-transaction-stats-merchant-yearly-failed", statsHandler.FindTransactionYearlyFailedByMerchant))
			byMerchant.Get("/monthly-method-success", params.apiHandler.Handle("get-transaction-stats-merchant-monthly-method-success", statsHandler.FindTransactionMonthlyMethodSuccessByMerchant))
			byMerchant.Get("/yearly-method-success", params.apiHandler.Handle("get-transaction-stats-merchant-yearly-method-success", statsHandler.FindTransactionYearlyMethodSuccessByMerchant))
			byMerchant.Get("/monthly-method-failed", params.apiHandler.Handle("get-transaction-stats-merchant-monthly-method-failed", statsHandler.FindTransactionMonthlyMethodFailedByMerchant))
			byMerchant.Get("/yearly-method-failed", params.apiHandler.Handle("get-transaction-stats-merchant-yearly-method-failed", statsHandler.FindTransactionYearlyMethodFailedByMerchant))
		})
	})

	return statsHandler
}

// parseYear reads and validates the "year" query parameter.
func parseYear(r *http.Request) (int, error) {
	year, err := strconv.Atoi(r.URL.Query().Get("year"))
	if err != nil || year < 1 || year > 9999 {
		return 0, sharedErrors.NewBadRequestError("invalid year")
	}
	return year, nil
}

// parseMonth reads and validates the "month" query parameter.
func parseMonth(r *http.Request) (int, error) {
	month, err := strconv.Atoi(r.URL.Query().Get("month"))
	if err != nil || month < 1 || month > 12 {
		return 0, sharedErrors.NewBadRequestError("invalid month")
	}
	return month, nil
}

// parseYearMonth reads both "year" and "month" query parameters.
func parseYearMonth(r *http.Request) (int, int, error) {
	year, err := parseYear(r)
	if err != nil {
		return 0, 0, err
	}
	month, err := parseMonth(r)
	if err != nil {
		return 0, 0, err
	}
	return year, month, nil
}

// parseMerchantID reads and validates the "merchant_id" query parameter.
func parseMerchantID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.URL.Query().Get("merchant_id"))
	if err != nil || id <= 0 {
		return 0, sharedErrors.NewBadRequestError("invalid merchant_id")
	}
	return id, nil
}

// parseCategoryID reads and validates the "category_id" query parameter.
func parseCategoryID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil || id <= 0 {
		return 0, sharedErrors.NewBadRequestError("invalid category_id")
	}
	return id, nil
}

func statsCacheKey(prefix string, parts ...int) string {
	key := prefix
	for _, p := range parts {
		key += fmt.Sprintf(":%d", p)
	}
	return key
}
