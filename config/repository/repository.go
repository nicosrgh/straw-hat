package repository

import (
	"fmt"
)

// Store :
type Store struct {
	// Mo *MongoDb
	My *MysqlDb
}

// Init :
func Init() *Store {
	fmt.Println("Starting server...")
	// mongo, err := InitMongo()
	// if err != nil {
	// 	logger.Error(err.Error(), err)
	// }

	// mysql, err := InitMysql()
	// if err != nil {
	// 	logger.Error(err.Error(), err)
	// }

	// mo := InitMongoStore(mongo)
	// my := InitMysqlStore(mysql)

	return &Store{
		// Mo: mo,
		// My: my,
	}
}
