package models

type PromoCode struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Discount        float64 `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type CreatePromoCode struct {
	Name            string  `json:"name"`
	Discount        float64 `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type PromoCodePrimaryKey struct {
	ID int `json:"id"`
}

type PromoCodeList struct {
	PromoCodes []*PromoCode `json:"promo_codes"`
}

type UpdatePromoCode struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Discount        float64 `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type GetListPromoCodeRequest struct {
	Search   string `json:"search"`
	OrderBy  string `json:"order_by"`
	OrderDir string `json:"order_dir"`
}

type GetListPromoCodeResponse struct {
	PromoCodes []*PromoCode `json:"promo_codes"`
}
