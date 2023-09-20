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

func (da *DataAccess) GetAccountData(pool *pgxpool.Pool, id string) (model.Account, error) {
	var account model.Account
	rows, err := pool.Query(context.Background(), "SELECT user_id, balance FROM data_balance WHERE user_id = $1", id)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println("error while iterating dataset")
		}
		account.ParseData(values)
		return account, nil
	}
	return account, pgx.ErrNoRows
}

func (da *DataAccess) AddAccountData(pool *pgxpool.Pool, id string) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO data_balance (user_id, balance) VALUES ($1, $2)", id, 0)
	if err != nil {
		return err
	}
	return err
}

func (da *DataAccess) IncreaseBalanceAccountData(pool *pgxpool.Pool, id string, amount int64) error {
	_, err := pool.Exec(context.Background(), "UPDATE data_balance SET balance = balance + $1 WHERE user_id = $2", amount, id)
	if err != nil {
		return err
	}
	return err
}

func (da *DataAccess) DecreaseBalanceAccountData(pool *pgxpool.Pool, id string, amount int64) error {
	_, err := pool.Exec(context.Background(), "UPDATE data_balance SET balance = balance - $1 WHERE user_id = $2", amount, id)
	if err != nil {
		return err
	}
	return err
}

func (da *DataAccess) GetHistoryOperationData(pool *pgxpool.Pool, id string) (model.HistoryOperation, error) {
	var history model.HistoryOperation
	rows, err := pool.Query(context.Background(), "SELECT * FROM history_table WHERE user_id = $1", id)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println("error while iterating dataset")
		}
		history.ParseData(values)
		return history, nil
	}
	return history, pgx.ErrNoRows
}

func (da *DataAccess) AddHistoryOperationData(pool *pgxpool.Pool, reqData model.DataRequest, operation string) error {
	var err error
	if operation != string(model.Transfer) {
		_, err = pool.Exec(context.Background(), "INSERT INTO history_table (user_id, operation, amount, to_id) VALUES ($1, $2, $3, $4)", reqData.UserId, operation, reqData.Amount, reqData.ToID)
		if err != nil {
			return err
		}
	} else {
		_, err = pool.Exec(context.Background(), "INSERT INTO history_table (user_id, operation, amount, to_id) VALUES ($1, $2, $3, $4)", reqData.FromID, operation, reqData.Amount, reqData.ToID)
		if err != nil {
			return err
		}
	}
	return err
}
