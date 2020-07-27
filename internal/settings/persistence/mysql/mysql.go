package mysql

import (
	"fmt"
	"net/url"

	grom "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlPersistence struct {
	db *grom.DB
}

func NewMysqlPersistence(dbConn string) (*MysqlPersistence, error) {
	db, err := grom.Open("mysql", dbConn)
	if err != nil {
		return &MysqlPersistence{}, err
	}

	err = db.DB().Ping()
	if err != nil {
		return &MysqlPersistence{}, err
	}

	return &MysqlPersistence{
		db: db,
	}, nil
}

// !!!: this is untested
func NewMySqlPersistenceWithURL(url *url.URL) (*MysqlPersistence, error) {
	dbUser := url.User.Username()
	dbPass, _ := url.User.Password()

	dbName := url.Path[1:len(url.Path)]
	if dbName == "" {
		dbName = "test"
	}

	dbConn := fmt.Sprintf("%s:%s@%s(%s)/%s", dbUser, dbPass, url.Scheme, url.Host, dbName)

	db, err := grom.Open("mysql", dbConn)
	if err != nil {
		return &MysqlPersistence{}, err
	}

	err = db.DB().Ping()
	if err != nil {
		return &MysqlPersistence{}, err
	}

	return &MysqlPersistence{
		db: db,
	}, nil
}
