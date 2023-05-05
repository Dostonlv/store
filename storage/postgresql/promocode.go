package postgresql

import (
	"app/api/models"
	_ "app/pkg/helper"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type promocodeRepo struct {
	db *pgxpool.Pool
}

func (p *promocodeRepo) Create(ctx context.Context, req *models.CreatePromoCode) (int, error) {
	var id int
	query := `
		INSERT INTO promo_codes(name,discount,discount_type,order_limit_price) 
		VALUES ($1,$2,$3,$4) RETURNING id
`
	err := p.db.QueryRow(ctx, query, req.Name, req.Discount, req.DiscountType, req.OrderLimitPrice).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *promocodeRepo) GetByID(ctx context.Context, req *models.PromoCodePrimaryKey) (*models.PromoCode, error) {

	promocode := &models.PromoCode{}
	query := `
		SELECT id,name,discount,discount_type,order_limit_price FROM promo_codes WHERE id = $1
	`
	err := p.db.QueryRow(ctx, query, req.ID).Scan(&promocode.ID, &promocode.Name, &promocode.Discount, &promocode.DiscountType, &promocode.OrderLimitPrice)
	if err != nil {
		return nil, err
	}
	return promocode, nil
}

func (p *promocodeRepo) GetList(ctx context.Context, req *models.GetListPromoCodeRequest) (resp *models.GetListPromoCodeResponse, err error) {
	resp = &models.GetListPromoCodeResponse{}
	var (
		query string
		rows  pgx.Rows
	)
	query = `
		SELECT id,name,discount,discount_type,order_limit_price FROM promo_codes
	`
	rows, err = p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		promocode := &models.PromoCode{}
		err = rows.Scan(&promocode.ID, &promocode.Name, &promocode.Discount, &promocode.DiscountType, &promocode.OrderLimitPrice)
		if err != nil {
			return nil, err
		}
		resp.PromoCodes = append(resp.PromoCodes, promocode)
	}
	return resp, nil
}

//func (p *promocodeRepo) Update(ctx context.Context, req *models.UpdatePromoCode) (int64, error) {
//	params := make(map[string]interface{})
//	query := `
//		UPDATE promo_codes SET
//		                    name = :name,
//		                    discount = :discount,
//		                    discount_type = :discount_type,
//		                    order_limit_price = :order_limit_price
// WHERE id = :id
//	`
//	params["id"] = req.ID
//	params["name"] = req.Name
//	params["discount"] = req.Discount
//	params["discount_type"] = req.DiscountType
//	params["order_limit_price"] = req.OrderLimitPrice
//	query, args := helper.ReplaceQueryParams(query, params)
//	result, err := p.db.Exec(ctx, query, args...)
//	if err != nil {
//		return 0, err
//	}
//
//	return result.RowsAffected(), nil
//
//}
//
//func (p promocodeRepo) Delete(ctx context.Context, req *models.PromoCodePrimaryKey) (int64, error) {
//	return 0, nil
//}
//
//func NewPromoCodeRepo(db *pgxpool.Pool) *promocodeRepo {
//	return &promocodeRepo{
//		db: db,
//	}
//}

func (p *promocodeRepo) Delete(ctx context.Context, req *models.PromoCodePrimaryKey) (int64, error) {
	query := `
		DELETE FROM
		 promo_codes 
		WHERE id = $1
		`

	result, err := p.db.Exec(ctx, query, req.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
func NewPromoCodeRepo(db *pgxpool.Pool) *promocodeRepo {
	return &promocodeRepo{
		db: db,
	}
}
