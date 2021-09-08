/*
@time: 2021/9/6 10:56
@author: chenZouLu
@file: mian_test
@software: GoLand
@note:
*/

package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB // 放在全局便于其他测试使用

func TestMain(m *testing.M) {
	var err error
	// func Open(driverName, dataSourceName string) (*DB, error) {}
	testDB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln("Connect err", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
