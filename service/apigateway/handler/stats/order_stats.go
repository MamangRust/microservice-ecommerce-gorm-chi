package statshandler

import (
	"net/http"

	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/httpx"
	pborder "github.com/MamangRust/microservice-ecommerce-grpc-pb/order"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/response"
	sharedErrors "github.com/MamangRust/microservice-ecommerce-shared/errors"

	gateway_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache/gateway_cache"
)

// @Security Bearer
// @Summary Monthly total revenue (order)
// @Tags Order Stats
// @Description Retrieve total order revenue for a given month from ClickHouse.
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseOrderMonthlyTotalRevenue
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/order/monthly-total-revenue [get]
func (h *statsHandleApi) FindOrderMonthlyTotalRevenue(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:order-stats:monthly-total-revenue", year, month)
	if cached, found := gateway_cache.Get[response.ApiResponseOrderMonthlyTotalRevenue](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.orderClient.FindMonthlyTotalRevenue(ctx, &pborder.FindYearMonthTotalRevenue{Year: int32(year), Month: int32(month)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseOrderMonthlyTotalRevenue(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly total revenue (order)
// @Tags Order Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponseOrderYearlyTotalRevenue
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/order/yearly-total-revenue [get]
func (h *statsHandleApi) FindOrderYearlyTotalRevenue(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:order-stats:yearly-total-revenue", year)
	if cached, found := gateway_cache.Get[response.ApiResponseOrderYearlyTotalRevenue](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.orderClient.FindYearlyTotalRevenue(ctx, &pborder.FindYearTotalRevenue{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseOrderYearlyTotalRevenue(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly revenue breakdown (order)
// @Tags Order Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponseOrderMonthly
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/order/monthly-revenue [get]
func (h *statsHandleApi) FindOrderMonthlyRevenue(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:order-stats:monthly-revenue", year)
	if cached, found := gateway_cache.Get[response.ApiResponseOrderMonthly](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.orderClient.FindMonthlyRevenue(ctx, &pborder.FindYearOrder{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseOrderMonthly(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly revenue breakdown (order)
// @Tags Order Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponseOrderYearly
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/order/yearly-revenue [get]
func (h *statsHandleApi) FindOrderYearlyRevenue(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:order-stats:yearly-revenue", year)
	if cached, found := gateway_cache.Get[response.ApiResponseOrderYearly](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.orderClient.FindYearlyRevenue(ctx, &pborder.FindYearOrder{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseOrderYearly(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly total revenue by merchant (order)
// @Tags Order Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseOrderMonthlyTotalRevenue
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/order/merchant/monthly-total-revenue [get]
func (h *statsHandleApi) FindOrderMonthlyTotalRevenueByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:order-stats:merchant:monthly-total-revenue", merchantID, year, month)
	if cached, found := gateway_cache.Get[response.ApiResponseOrderMonthlyTotalRevenue](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.orderByMerchantClient.FindMonthlyTotalRevenueByMerchant(ctx, &pborder.FindYearMonthTotalRevenueByMerchant{
		Year: int32(year), Month: int32(month), MerchantId: int32(merchantID),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseOrderMonthlyTotalRevenue(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly total revenue by merchant (order)
// @Tags Order Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseOrderYearlyTotalRevenue
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/order/merchant/yearly-total-revenue [get]
func (h *statsHandleApi) FindOrderYearlyTotalRevenueByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:order-stats:merchant:yearly-total-revenue", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponseOrderYearlyTotalRevenue](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.orderByMerchantClient.FindYearlyTotalRevenueByMerchant(ctx, &pborder.FindYearTotalRevenueByMerchant{
		Year: int32(year), MerchantId: int32(merchantID),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseOrderYearlyTotalRevenue(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly revenue breakdown by merchant (order)
// @Tags Order Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseOrderMonthly
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/order/merchant/monthly-revenue [get]
func (h *statsHandleApi) FindOrderMonthlyRevenueByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:order-stats:merchant:monthly-revenue", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponseOrderMonthly](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.orderByMerchantClient.FindMonthlyRevenueByMerchant(ctx, &pborder.FindYearOrderByMerchant{
		MerchantId: int32(merchantID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseOrderMonthly(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly revenue breakdown by merchant (order)
// @Tags Order Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseOrderYearly
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/order/merchant/yearly-revenue [get]
func (h *statsHandleApi) FindOrderYearlyRevenueByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:order-stats:merchant:yearly-revenue", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponseOrderYearly](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.orderByMerchantClient.FindYearlyRevenueByMerchant(ctx, &pborder.FindYearOrderByMerchant{
		MerchantId: int32(merchantID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseOrderYearly(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}
