package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/nicosrgh/straw-hat/lib/logger"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMysql :
func InitMysql(conn string, name string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s%s", conn, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return db, nil
}

// MysqlDb :
type MysqlDb struct {
	*sql.DB
}

// MysqlStore :
type MysqlStore interface {
	Read(filter interface{}, opts options.FindOneOptions, collectionName string) *mongo.SingleResult
}

// Read :
func (my *MysqlDb) Read(query string) (string, error) {
	rows, err := my.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// InitMysqlStore :
func InitMysqlStore(db *sql.DB) *MysqlDb {
	return &MysqlDb{db}
}
