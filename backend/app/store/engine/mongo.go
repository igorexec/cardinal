package engine

import (
	"context"
	"github.com/igorexec/cardinal/app/store"
	"go.mongodb.org/mongo-driver/bson"
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

func NewMongo(uri string) (*Mongo, error) {
	ctx := context.Background()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &Mongo{
		db: db.Database(database),
	}, nil
}

func (m *Mongo) Create(ctx context.Context, pageSpeed store.PageSpeed) (pageSpeedID string, err error) {
	c := m.db.Collection(mongoPageSpeed)
	result, err := c.InsertOne(ctx, pageSpeed)
	if err != nil {
		log.Printf("[error] failed to pagespeed insert to DB: %v", err)
		return "", err
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (m *Mongo) Get(ctx context.Context, page string, from time.Time, to time.Time) (ps []store.PageSpeed, err error) {
	c := m.db.Collection(mongoPageSpeed)
	cur, err := c.Find(ctx, bson.M{"date": bson.M{"$gt": from, "$lt": to}, "page": page})
	if err != nil {
		log.Printf("[warn] pagespeed from %s to %s not found", from.String(), to.String())
		return nil, err
	}

	for cur.Next(ctx) {
		var elem store.PageSpeed
		if err := cur.Decode(&elem); err != nil {
			log.Fatalf("[error] failed to decode: %v", err)
		}

		ps = append(ps, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatalf("[error] failed to iterate: %v", err)
	}
	if err := cur.Close(ctx); err != nil {
		log.Fatalf("[error] failed to close cursor: %v", err)
	}
	return ps, nil
}

func (m *Mongo) Close(ctx context.Context) error {
	if err := m.client.Disconnect(ctx); err != nil {
		log.Fatalf("[error] failed to close MongoDB connection: %v", err)
		return err
	}
	return nil
}
