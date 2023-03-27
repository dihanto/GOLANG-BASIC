package main

import (
	"context"
	"time"

	"github.com/dihanto/golang-basic/17-database-mysql/config"
	"github.com/dihanto/golang-basic/17-database-mysql/entity"
	"github.com/dihanto/golang-basic/17-database-mysql/repository/mysql"
	"github.com/dihanto/golang-basic/17-database-mysql/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.NewDB()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	productRepository := mysql.NewMysqlProductRepostitory(db)

	timeoutContext := time.Duration(2) * time.Second

	productUsecase := usecase.NewProductUsecase(productRepository, timeoutContext)
	ctx := context.Background()
	product := entity.Product{
		Name:  "odading",
		Price: 2000,
	}
	productUsecase.Insert(ctx, &product)
}
