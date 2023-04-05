package mysql

import (
	"context"
	"database/sql"
	"log"

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

func (m *mysqlProductRepository) FindAll(ctx context.Context) (products []entity.Product, err error) {
	query := `SELECT  id, name, price FROM products`
	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println(errRow)
		}
	}()
	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			log.Println(err)
			return
		}
		products = append(products, product)
	}
	return
}

func (m *mysqlProductRepository) Update(ctx context.Context, product *entity.Product) (err error) {
	query := `UPDATE products SET name=?, price=? WHERE id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, product.Name, product.Price, product.ID)
	if err != nil {
		return
	}

	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	return
	// }

	// if rowsAffected != 1 {
	// 	err = fmt.Errorf("weird behavior. rows affected : %d", rowsAffected)
	// 	return
	// }
	return
}

func (m *mysqlProductRepository) Delete(ctx context.Context, id int) (err error) {
	query := `Delete from products WHERE id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	return
	// }

	// if rowsAffected != 1 {
	// 	err = fmt.Errorf("weird behavior. rows affected : %d", rowsAffected)
	// 	return
	// }
	return
}
