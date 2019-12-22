package engine

import (
	"context"
	"github.com/igorexec/cardinal/app/store"
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

	ctx context.Context
}

func NewMongo(uri string) *Mongo {
	ctx := context.Background()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return &Mongo{
		client: db,
		ctx:    ctx,
	}
}

func (m *Mongo) Create(pageSpeed store.PageSpeed) (pageSpeedID string, err error) {
	return "", nil
}

func (m *Mongo) Get(from time.Time, to time.Time) ([]store.PageSpeed, error) {
	return nil, nil
}

func (m *Mongo) Close() error {
	if err := m.client.Disconnect(m.ctx); err != nil {
		log.Fatalf("[error] failed to close MongoDB connection: %v", err)
		return err
	}
	return nil
}
