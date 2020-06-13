package repository

import (
	"context"
	"time"

	"github.com/nicosrgh/straw-hat/config"
	"github.com/nicosrgh/straw-hat/lib/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongo :
func InitMongo() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.C.MongoDbDsn))
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	db := client.Database(config.C.MongoDbName)

	return db, nil
}

// MongoDb :
type MongoDb struct {
	*mongo.Database
}

// MongoStore :
type MongoStore interface {
	Aggregate(pipeline interface{}, collectionName string) (mongo.Cursor, error)
	Create(document interface{}, collectionName string) (*mongo.InsertOneResult, error)
	Read(filter interface{}, opts options.FindOneOptions, collectionName string) *mongo.SingleResult
	ReadAll(filter interface{}, opts options.FindOptions, collectionName string) (mongo.Cursor, error)
	Update(document interface{}, filter interface{}, collectionName string) (*mongo.UpdateResult, error)
	UpdateAll(documents []interface{}, filter interface{}, collectionName string) (*mongo.UpdateResult, error)
}

// MongoDocument :
type MongoDocument struct {
	Data interface{}
}

// Aggregate :
func (m *MongoDb) Aggregate(pipeline interface{}, collectionName string) (*mongo.Cursor, error) {
	col := m.Collection(collectionName)
	return col.Aggregate(context.Background(), pipeline)
}

// Create :
func (m *MongoDb) Create(document interface{}, collectionName string) (*mongo.InsertOneResult, error) {
	col := m.Collection(collectionName)
	return col.InsertOne(context.Background(), document)
}

// Read :
func (m *MongoDb) Read(filter interface{}, opts options.FindOneOptions, collectionName string) *mongo.SingleResult {
	col := m.Collection(collectionName)
	return col.FindOne(context.Background(), filter, &opts)
}

// ReadAll :
func (m *MongoDb) ReadAll(filter interface{}, opts options.FindOptions, collectionName string) (*mongo.Cursor, error) {
	col := m.Collection(collectionName)
	return col.Find(context.Background(), filter, &opts)
}

// Update :
func (m *MongoDb) Update(filter interface{}, document interface{}, collectionName string) (*mongo.UpdateResult, error) {
	col := m.Collection(collectionName)
	return col.UpdateOne(context.Background(), filter, document)
}

// UpdateAll :
func (m *MongoDb) UpdateAll(filter interface{}, documents []interface{}, collectionName string) (*mongo.UpdateResult, error) {
	col := m.Collection(collectionName)
	return col.UpdateMany(context.Background(), filter, documents)
}

// Delete :
func (m *MongoDb) Delete(filter interface{}, collectionName string) (*mongo.DeleteResult, error) {
	col := m.Collection(collectionName)
	return col.DeleteOne(context.Background(), filter)
}

// DeleteAll :
func (m *MongoDb) DeleteAll(filter interface{}, collectionName string) (*mongo.DeleteResult, error) {
	col := m.Collection(collectionName)
	return col.DeleteMany(context.Background(), filter)
}

// InitMongoStore :
func InitMongoStore(db *mongo.Database) *MongoDb {
	return &MongoDb{db}
}
