package stock

import (
	"context"

	entity "github.com/wahyurudiyan/warehouse-stock/entity/stock"
	"github.com/wahyurudiyan/warehouse-stock/repository/stock"
)

type usecase struct {
	repo stock.Repository
}

type Usecase interface {
	CreateStock(ctx context.Context, s []*entity.Product) error
}

func NewUsecase(r stock.Repository) Usecase {
	return &usecase{r}
}

func (u *usecase) CreateStock(ctx context.Context, s []*entity.Product) error {
	return u.repo.CreateStock(ctx, s)
}
