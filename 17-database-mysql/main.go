package main

import (
	"context"
	"log"
	"time"

	"github.com/dihanto/golang-basic/17-database-mysql/config"
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
	// product := entity.Product{
	// 	ID:    2,
	// 	Name:  "dading",
	// 	Price: 3000,
	// }
	// productUsecase.Insert(ctx, &product)
	// products, err := productUsecase.FindAll(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(products)
	// err = productUsecase.Update(ctx, &product)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err = productUsecase.Delete(ctx, 4)
	if err != nil {
		log.Fatal(err)
	}
}
