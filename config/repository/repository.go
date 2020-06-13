package repository

import (
	"fmt"

	"github.com/nicosrgh/straw-hat/lib/logger"
)

// Store :
type Store struct {
	Mo *MongoDb
}

// Init :
func Init() *Store {
	fmt.Println("Starting server...")
	mongo, err := InitMongo()

	if err != nil {
		logger.Error(err.Error(), err)
	}

	mo := InitMongoStore(mongo)

	return &Store{
		Mo: mo,
	}
}
