package statshandler

import (
	"net/http"

	gateway_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache/gateway_cache"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/httpx"
	pbcategory "github.com/MamangRust/microservice-ecommerce-grpc-pb/category"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/response"
	sharedErrors "github.com/MamangRust/microservice-ecommerce-shared/errors"
)

// @Security Bearer
// @Summary Monthly total pricing (category)
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseCategoryMonthlyTotalPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/monthly-total-pricing [get]
func (h *statsHandleApi) FindCategoryMonthlyTotalPrice(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:monthly-total-pricing", year, month)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryMonthlyTotalPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryClient.FindMonthlyTotalPrices(ctx, &pbcategory.FindYearMonthTotalPrices{Year: int32(year), Month: int32(month)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryMonthlyTotalPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly total pricing (category)
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponseCategoryYearlyTotalPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/yearly-total-pricing [get]
func (h *statsHandleApi) FindCategoryYearlyTotalPrice(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:yearly-total-pricing", year)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryYearlyTotalPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryClient.FindYearlyTotalPrices(ctx, &pbcategory.FindYearTotalPrices{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryYearlyTotalPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly pricing breakdown (category)
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponseCategoryMonthPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/monthly-pricing [get]
func (h *statsHandleApi) FindCategoryMonthPrice(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:monthly-pricing", year)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryMonthPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryClient.FindMonthPrice(ctx, &pbcategory.FindYearCategory{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryMonthPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly pricing breakdown (category)
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Success 200 {object} response.ApiResponseCategoryYearPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/yearly-pricing [get]
func (h *statsHandleApi) FindCategoryYearPrice(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:yearly-pricing", year)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryYearPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryClient.FindYearPrice(ctx, &pbcategory.FindYearCategory{Year: int32(year)})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryYearPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly total pricing by category id
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Param category_id query int true "Category ID"
// @Success 200 {object} response.ApiResponseCategoryMonthlyTotalPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/mycategory/monthly-total-pricing [get]
func (h *statsHandleApi) FindCategoryMonthlyTotalPriceById(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}
	categoryID, err := parseCategoryID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:mycategory:monthly-total-pricing", categoryID, year, month)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryMonthlyTotalPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryByIdClient.FindMonthlyTotalPricesById(ctx, &pbcategory.FindYearMonthTotalPriceById{
		Year: int32(year), Month: int32(month), CategoryId: int32(categoryID),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryMonthlyTotalPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly total pricing by category id
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param category_id query int true "Category ID"
// @Success 200 {object} response.ApiResponseCategoryYearlyTotalPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/mycategory/yearly-total-pricing [get]
func (h *statsHandleApi) FindCategoryYearlyTotalPriceById(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	categoryID, err := parseCategoryID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:mycategory:yearly-total-pricing", categoryID, year)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryYearlyTotalPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryByIdClient.FindYearlyTotalPricesById(ctx, &pbcategory.FindYearTotalPriceById{
		Year: int32(year), CategoryId: int32(categoryID),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryYearlyTotalPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly pricing breakdown by category id
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param category_id query int true "Category ID"
// @Success 200 {object} response.ApiResponseCategoryMonthPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/mycategory/monthly-pricing [get]
func (h *statsHandleApi) FindCategoryMonthPriceById(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	categoryID, err := parseCategoryID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:mycategory:monthly-pricing", categoryID, year)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryMonthPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryByIdClient.FindMonthPriceById(ctx, &pbcategory.FindYearCategoryById{
		CategoryId: int32(categoryID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryMonthPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly pricing breakdown by category id
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param category_id query int true "Category ID"
// @Success 200 {object} response.ApiResponseCategoryYearPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/mycategory/yearly-pricing [get]
func (h *statsHandleApi) FindCategoryYearPriceById(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	categoryID, err := parseCategoryID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:mycategory:yearly-pricing", categoryID, year)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryYearPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryByIdClient.FindYearPriceById(ctx, &pbcategory.FindYearCategoryById{
		CategoryId: int32(categoryID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryYearPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly total pricing by merchant
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param month query int true "Month (1-12)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseCategoryMonthlyTotalPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/merchant/monthly-total-pricing [get]
func (h *statsHandleApi) FindCategoryMonthlyTotalPriceByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, month, err := parseYearMonth(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:merchant:monthly-total-pricing", merchantID, year, month)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryMonthlyTotalPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryByMerchantCli.FindMonthlyTotalPricesByMerchant(ctx, &pbcategory.FindYearMonthTotalPriceByMerchant{
		Year: int32(year), Month: int32(month), MerchantId: int32(merchantID),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryMonthlyTotalPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly total pricing by merchant
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseCategoryYearlyTotalPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/merchant/yearly-total-pricing [get]
func (h *statsHandleApi) FindCategoryYearlyTotalPriceByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:merchant:yearly-total-pricing", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryYearlyTotalPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryByMerchantCli.FindYearlyTotalPricesByMerchant(ctx, &pbcategory.FindYearTotalPriceByMerchant{
		Year: int32(year), MerchantId: int32(merchantID),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryYearlyTotalPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Monthly pricing breakdown by merchant
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseCategoryMonthPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/merchant/monthly-pricing [get]
func (h *statsHandleApi) FindCategoryMonthPriceByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:merchant:monthly-pricing", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryMonthPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryByMerchantCli.FindMonthPriceByMerchant(ctx, &pbcategory.FindYearCategoryByMerchant{
		MerchantId: int32(merchantID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryMonthPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}

// @Security Bearer
// @Summary Yearly pricing breakdown by merchant
// @Tags Category Stats
// @Accept json
// @Produce json
// @Param year query int true "Year (e.g. 2026)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseCategoryYearPrice
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/category-stats/merchant/yearly-pricing [get]
func (h *statsHandleApi) FindCategoryYearPriceByMerchant(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	year, err := parseYear(r)
	if err != nil {
		return err
	}
	merchantID, err := parseMerchantID(r)
	if err != nil {
		return err
	}

	cacheKey := statsCacheKey("apigw:category-stats:merchant:yearly-pricing", merchantID, year)
	if cached, found := gateway_cache.Get[response.ApiResponseCategoryYearPrice](ctx, h.cache, cacheKey); found && cached != nil {
		return httpx.JSON(w, http.StatusOK, cached)
	}

	res, err := h.categoryByMerchantCli.FindYearPriceByMerchant(ctx, &pbcategory.FindYearCategoryByMerchant{
		MerchantId: int32(merchantID), Year: int32(year),
	})
	if err != nil {
		return sharedErrors.ParseGrpcError(err)
	}

	so := h.mapping.ToApiResponseCategoryYearPrice(res)
	gateway_cache.Set(ctx, h.cache, cacheKey, so, statsCacheTTL)
	return httpx.JSON(w, http.StatusOK, so)
}
