package postgresql

import (
	"app/api/models"
	"context"
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

func (p promocodeRepo) GetByID(ctx context.Context, req *models.PromoCodePrimaryKey) (*models.PromoCode, error) {
	return nil, nil
}

func (p promocodeRepo) GetList(ctx context.Context, req *models.GetListPromoCodeRequest) (resp *models.GetListPromoCodeResponse, err error) {
	return resp, nil
}

func (p promocodeRepo) Update(ctx context.Context, req *models.UpdatePromoCode) (int64, error) {
	return 0, nil
}

func (p promocodeRepo) Delete(ctx context.Context, req *models.PromoCodePrimaryKey) (int64, error) {
	return 0, nil
}

func NewPromoCodeRepo(db *pgxpool.Pool) *promocodeRepo {
	return &promocodeRepo{
		db: db,
	}
}
