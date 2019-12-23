package engine

import (
	"context"
	"github.com/igorexec/cardinal/app/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	mongoPageSpeed = "pagespeed"
)

type Mongo struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongo(uri string) *Mongo {
	ctx := context.Background()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return &Mongo{
		db: db.Database(database),
	}
}

func (m *Mongo) Create(ctx context.Context, pageSpeed store.PageSpeed) (pageSpeedID string, err error) {
	collection := m.db.Collection(mongoPageSpeed)
	result, err := collection.InsertOne(ctx, pageSpeed)
	if err != nil {
		log.Printf("[error] failed to pagespeed insert to DB: %v", err)
		return "", err
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (m *Mongo) Get(ctx context.Context, from time.Time, to time.Time) ([]store.PageSpeed, error) {
	return nil, nil
}

func (m *Mongo) Close(ctx context.Context) error {
	if err := m.client.Disconnect(ctx); err != nil {
		log.Fatalf("[error] failed to close MongoDB connection: %v", err)
		return err
	}
	return nil
}
