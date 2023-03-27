package mysql

import (
	"context"
	"database/sql"

	"github.com/dihanto/golang-basic/17-database-mysql/entity"
)

type mysqlProductRepository struct {
	Conn *sql.DB
}

func NewMysqlProductRepostitory(conn *sql.DB) entity.ProductRepository {
	return &mysqlProductRepository{conn}
}

func (m *mysqlProductRepository) Insert(ctx context.Context, product *entity.Product) (err error) {
	query := `INSERT products SET name=?, price=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, product.Name, product.Price)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	product.ID = lastID
	return
}

func (m *mysqlProductRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	panic("not implemented") // TODO: Implement
}
