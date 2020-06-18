package repository

import (
	"fmt"

	"github.com/nicosrgh/straw-hat/config"
	"github.com/nicosrgh/straw-hat/lib/logger"
)

// Store :
type Store struct {
	Mo *MongoDb
	My *MysqlDb
}

// Init :
func Init() *Store {
	fmt.Println("Starting server...")
	mongo, err := InitMongo()
	if err != nil {
		logger.Error(err.Error(), err)
	}

	mysql, err := InitMysql(config.C.MySqlDbDsn, config.C.MySqlDbName)
	if err != nil {
		logger.Error(err.Error(), err)
	}

	mo := InitMongoStore(mongo)
	my := InitMysqlStore(mysql)

	return &Store{
		Mo: mo,
		My: my,
	}
}
