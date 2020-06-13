package stock

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/appscode/go/strings"
	"github.com/uber/tchannel-go/crossdock/log"
	"github.com/wahyurudiyan/warehouse-stock/entity/stock"
)

type repository struct {
	conn *sql.DB
}

type Repository interface {
	CreateStock(ctx context.Context, s []*stock.Product) error
}

func NewRepository(conn *sql.DB) Repository {
	return &repository{conn}
}

func (r *repository) CreateStock(ctx context.Context, s []*stock.Product) error {
	var args []*string

	query := `INSERT INTO warehouse.stock
	(name, category, producer, price, quantity, production_date, expired_date, create_at, update_at)
	VALUES %s`

	for _, v := range s {
		arg := fmt.Sprintf("('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v')",
			v.Name, v.Category, v.Producer, v.Price, v.Quantity, v.ProductionDate, v.ExpiredDate, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))

		args = append(args, &arg)
	}

	a := strings.Join(args, ", ")
	q := fmt.Sprintf(query, a)

	log.Printf("%s", q)

	stmt, err := r.conn.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil

}
