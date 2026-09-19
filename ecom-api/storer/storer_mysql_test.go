package storer

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestCreateProduct(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	st := NewMySQLStorer(db)

	// create product struct
	p := &Product{
		Name:         "test product",
		Image:        "test.jpg",
		Category:     "test category",
		Description:  "test description",
		Rating:       5,
		NumReviews:   10,
		Price:        100.0,
		CountInStock: 100,
	}

	mock.ExpectExec("INSERT INTO products (name, image, category, description, rating, num_reviews, price, count_in_stock) VALUES (?, ?, ?, ?,?, ?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
	_, err = st.CreateProduct(context.Background(), p)
	if err != nil {
		t.Fatalf("error creating product: %v", err)
	}
}
