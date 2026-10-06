package statshandler

import (
	"net/http"

	gateway_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache/gateway_cache"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/httpx"
	pbtransaction "github.com/MamangRust/microservice-ecommerce-grpc-pb/transaction"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/response"
	sharedErrors "github.com/MamangRust/microservice-ecommerce-shared/errors"
)

// @Security Bearer
// @Summary Monthly successful transaction amount
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthSuccess
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/monthly-success [get]
func (h *statsHandleApi) FindTransactionMonthlySuccess(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:monthly-success", year, month)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionMonthSuccess](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionClient.GetMonthlyAmountSuccess(ctx, &pbtransaction.MonthAmountTransactionRequest{Year: int32(year), Month: int32(month)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionMonthSuccess(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly successful transaction amount
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponsesTransactionYearSuccess
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/yearly-success [get]
func (h *statsHandleApi) FindTransactionYearlySuccess(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:yearly-success", year)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionYearSuccess](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionClient.GetYearlyAmountSuccess(ctx, &pbtransaction.YearAmountTransactionRequest{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionYearSuccess(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly failed transaction amount
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthFailed
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/monthly-failed [get]
func (h *statsHandleApi) FindTransactionMonthlyFailed(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:monthly-failed", year, month)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionMonthFailed](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionClient.GetMonthlyAmountFailed(ctx, &pbtransaction.MonthAmountTransactionRequest{Year: int32(year), Month: int32(month)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionMonthFailed(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly failed transaction amount
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponsesTransactionYearFailed
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/yearly-failed [get]
func (h *statsHandleApi) FindTransactionYearlyFailed(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:yearly-failed", year)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionYearFailed](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionClient.GetYearlyAmountFailed(ctx, &pbtransaction.YearAmountTransactionRequest{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionYearFailed(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly payment method breakdown (success)
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/monthly-method-success [get]
func (h *statsHandleApi) FindTransactionMonthlyMethodSuccess(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:monthly-method-success", year, month)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionMonthMethod](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionClient.GetMonthlyTransactionMethodSuccess(ctx, &pbtransaction.MonthMethodTransactionRequest{Year: int32(year), Month: int32(month)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionMonthMethod(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly payment method breakdown (success)
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/yearly-method-success [get]
func (h *statsHandleApi) FindTransactionYearlyMethodSuccess(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:yearly-method-success", year)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionYearMethod](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionClient.GetYearlyTransactionMethodSuccess(ctx, &pbtransaction.YearMethodTransactionRequest{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionYearMethod(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly payment method breakdown (failed)
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/monthly-method-failed [get]
func (h *statsHandleApi) FindTransactionMonthlyMethodFailed(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:monthly-method-failed", year, month)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionMonthMethod](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionClient.GetMonthlyTransactionMethodFailed(ctx, &pbtransaction.MonthMethodTransactionRequest{Year: int32(year), Month: int32(month)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionMonthMethod(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly payment method breakdown (failed)
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/yearly-method-failed [get]
func (h *statsHandleApi) FindTransactionYearlyMethodFailed(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:yearly-method-failed", year)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionYearMethod](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionClient.GetYearlyTransactionMethodFailed(ctx, &pbtransaction.YearMethodTransactionRequest{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionYearMethod(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly successful transaction amount by merchant
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponsesTransactionMonthSuccess
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/merchant/monthly-success [get]
func (h *statsHandleApi) FindTransactionMonthlySuccessByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:merchant:monthly-success", merchantID, year, month)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionMonthSuccess](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionByMerchant.GetMonthlyAmountSuccessByMerchant(ctx, &pbtransaction.MonthAmountTransactionMerchantRequest{
		MerchantId: int32(merchantID), Year: int32(year), Month: int32(month),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionMonthSuccess(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly successful transaction amount by merchant
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponsesTransactionYearSuccess
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/merchant/yearly-success [get]
func (h *statsHandleApi) FindTransactionYearlySuccessByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:merchant:yearly-success", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionYearSuccess](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionByMerchant.GetYearlyAmountSuccessByMerchant(ctx, &pbtransaction.YearAmountTransactionMerchantRequest{
		MerchantId: int32(merchantID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionYearSuccess(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly failed transaction amount by merchant
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponsesTransactionMonthFailed
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/merchant/monthly-failed [get]
func (h *statsHandleApi) FindTransactionMonthlyFailedByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:merchant:monthly-failed", merchantID, year, month)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionMonthFailed](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionByMerchant.GetMonthlyAmountFailedByMerchant(ctx, &pbtransaction.MonthAmountTransactionMerchantRequest{
		MerchantId: int32(merchantID), Year: int32(year), Month: int32(month),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionMonthFailed(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly failed transaction amount by merchant
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponsesTransactionYearFailed
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/merchant/yearly-failed [get]
func (h *statsHandleApi) FindTransactionYearlyFailedByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:merchant:yearly-failed", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionYearFailed](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionByMerchant.GetYearlyAmountFailedByMerchant(ctx, &pbtransaction.YearAmountTransactionMerchantRequest{
		MerchantId: int32(merchantID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionYearFailed(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly payment method breakdown by merchant (success)
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/merchant/monthly-method-success [get]
func (h *statsHandleApi) FindTransactionMonthlyMethodSuccessByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:merchant:monthly-method-success", merchantID, year, month)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionMonthMethod](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionByMerchant.GetMonthlyTransactionMethodByMerchantSuccess(ctx, &pbtransaction.MonthMethodTransactionMerchantRequest{
		MerchantId: int32(merchantID), Year: int32(year), Month: int32(month),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionMonthMethod(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly payment method breakdown by merchant (success)
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/merchant/yearly-method-success [get]
func (h *statsHandleApi) FindTransactionYearlyMethodSuccessByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:merchant:yearly-method-success", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionYearMethod](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionByMerchant.GetYearlyTransactionMethodByMerchantSuccess(ctx, &pbtransaction.YearMethodTransactionMerchantRequest{
		MerchantId: int32(merchantID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionYearMethod(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly payment method breakdown by merchant (failed)
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/merchant/monthly-method-failed [get]
func (h *statsHandleApi) FindTransactionMonthlyMethodFailedByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:merchant:monthly-method-failed", merchantID, year, month)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionMonthMethod](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionByMerchant.GetMonthlyTransactionMethodByMerchantFailed(ctx, &pbtransaction.MonthMethodTransactionMerchantRequest{
		MerchantId: int32(merchantID), Year: int32(year), Month: int32(month),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionMonthMethod(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly payment method breakdown by merchant (failed)
// @Tags Transaction Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/transaction-stats/merchant/yearly-method-failed [get]
func (h *statsHandleApi) FindTransactionYearlyMethodFailedByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:transaction-stats:merchant:yearly-method-failed", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponsesTransactionYearMethod](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.transactionByMerchant.GetYearlyTransactionMethodByMerchantFailed(ctx, &pbtransaction.YearMethodTransactionMerchantRequest{
		MerchantId: int32(merchantID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseTransactionYearMethod(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}
