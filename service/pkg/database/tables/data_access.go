package tables

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"service/pkg/model"
)

type DataAccess struct {
}

func (da *DataAccess) GetData(pool *pgxpool.Pool, id string) (model.Account, error) {
	var tableData model.Account
	rows, err := pool.Query(context.Background(), "SELECT user_id, balance FROM data_balance WHERE user_id = $1", id)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println("error while iterating dataset")
		}
		tableData.ParseData(values)
		return tableData, nil
	}
	return tableData, pgx.ErrNoRows
}

func (da *DataAccess) AddData(pool *pgxpool.Pool, id string) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO data_balance (user_id, balance) VALUES ($1, $2)", id, 0)
	if err != nil {
		return err
	}
	return err
}

func (da *DataAccess) IncreaseData(pool *pgxpool.Pool, id string, amount int64) error {
	_, err := pool.Query(context.Background(), "UPDATE data_balance SET balance = balance + $1 WHERE user_id = $2", amount, id)
	if err != nil {
		return err
	}
	return err
}

func (da *DataAccess) DecreaseData(pool *pgxpool.Pool, id string, amount int64) error {
	_, err := pool.Exec(context.Background(), "UPDATE data_balance SET balance = balance - $1 WHERE user_id = $2", amount, id)
	if err != nil {
		return err
	}
	return err
}
