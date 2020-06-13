package main

import (
	"github.com/labstack/echo"
	"github.com/wahyurudiyan/warehouse-stock/adapter/http"
	"github.com/wahyurudiyan/warehouse-stock/adapter/mysql"
	"github.com/wahyurudiyan/warehouse-stock/config"
	repositoryLayer "github.com/wahyurudiyan/warehouse-stock/repository/stock"
	usecaseLayer "github.com/wahyurudiyan/warehouse-stock/usecase/stock"
)

func main() {
	newEcho := echo.New()
	cfg := config.GetConfiguration()
	sql, err := mysql.NewConnection(cfg)
	if err != nil {
		panic(err)
	}

	repo := repositoryLayer.NewRepository(sql)
	usecase := usecaseLayer.NewUsecase(repo)
	serviceHandler := http.NewHandler(usecase)
	router := http.NewServiceRouter(serviceHandler)

	router.Router(newEcho)

	if err = newEcho.Start(cfg.PortHTTP); err != nil {
		panic(err)
	}
}
