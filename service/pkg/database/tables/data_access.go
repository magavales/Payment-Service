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

func (da *DataAccess) GetData(pool *pgxpool.Pool, dr model.DataRequest) (model.Account, error) {
	var tableData model.Account
	rows, err := pool.Query(context.Background(), "SELECT user_id, balance FROM data_balance WHERE user_id = $1", dr.User_ID)
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

func (da *DataAccess) IncreaseData(pool *pgxpool.Pool, dr model.DataRequest) error {
	_, err := pool.Exec(context.Background(), "UPDATE data_balance SET balance = balance + $1 WHERE user_id = $2", dr.Amount, dr.User_ID)
	if err != nil {
		return err
	}
	return err
}

func (da *DataAccess) DecreaseData(pool *pgxpool.Pool, dr model.DataRequest) error {
	_, err := pool.Exec(context.Background(), "UPDATE data_balance SET balance = balance - $1 WHERE user_id = $2", dr.Amount, dr.User_ID)
	if err != nil {
		return err
	}
	return err
}
