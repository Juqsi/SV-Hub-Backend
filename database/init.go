package database

import (
	"HexMaster/utils"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

const (
	DatabaseDateTimeFormat string = "2006-01-02 15:04:05"
)

type DB struct {
	*sql.DB
}

type Handler func(rows *sql.Rows) error

var (
	db             *DB = initDB()
	HandlerNothing     = func(rows *sql.Rows) error { return nil }
)

func initDB() *DB {
	var err error

	config := mysql.Config{
		User:                     utils.GetEnv("DATABASE_USER", "[backend_user]"),
		Passwd:                   utils.GetEnv("DATABASE_PASSWORD", "[backend_password]"),
		Net:                      "tcp",
		Addr:                     utils.GetEnv("DATABASE_IP", "localhost:3306"),
		DBName:                   utils.GetEnv("DATABASE_NAME", "SVHub"),
		Params:                   nil,
		Collation:                "",
		Loc:                      nil,
		MaxAllowedPacket:         0,
		ServerPubKey:             "",
		TLS:                      nil,
		Timeout:                  0,
		ReadTimeout:              0,
		WriteTimeout:             0,
		AllowAllFiles:            false,
		AllowCleartextPasswords:  false,
		AllowFallbackToPlaintext: false,
		AllowNativePasswords:     true,
		AllowOldPasswords:        false,
		CheckConnLiveness:        false,
		ClientFoundRows:          false,
		ColumnsWithAlias:         false,
		InterpolateParams:        false,
		MultiStatements:          false,
		ParseTime:                true,
		RejectReadOnly:           false,
	}

	var newDB *sql.DB
	err = fmt.Errorf("database not connected")

	for err != nil {
		err = nil
		newDB, err = sql.Open("mysql", config.FormatDSN())
		if err != nil {
			fmt.Println("❌  unable to open database connection -> " + err.Error())
			time.Sleep(10 * time.Second)
			continue
		}

		err = newDB.Ping()
		if err != nil {
			fmt.Println("❌  no database connection -> " + err.Error())
			time.Sleep(10 * time.Second)
			continue
		}
	}
	fmt.Println("✅  connected")
	return &DB{newDB}
}
