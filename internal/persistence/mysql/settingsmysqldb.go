package mysql

import (
	"database/sql"
	"fmt"
)

type SettingsMySqlDB struct {
	db *sql.DB
}

func NewMySqlDB(dbUser, dbPass, dbSocket, dbHost, dbName string) (*SettingsMySqlDB, error) {
	dbInfo := fmt.Sprintf("%s:%s@%s(%s)/%s")
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		return &SettingsMySqlDB{}, err
	}

	err = db.Ping()
	if err != nil {
		return &SettingsMySqlDB{}, err
	}

	return &SettingsMySqlDB{
		db: db,
	}, nil
}
