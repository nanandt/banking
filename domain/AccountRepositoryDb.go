package domain

import (
	"banking/errs"
	"banking/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values(?, ?, ?, ?, ?) "

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`insert into transactions (account_id, amount, transaction_type, transaction_date)
							values(?,?,?,?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`update accounts set amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`update accounts set amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be received
	if err != nil {
		tx.Rollback()
		logger.Error("error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// commit transaction when  all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// getting the last transaction ID from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected  database error")
	}

	// getting the latest account infomation from the accounts table
	account, appError := d.FindBy(t.AccountId)
	if appError != nil {
		return nil, appError
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the last balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "select account_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewValidationError("unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
