package engine

import "go.mongodb.org/mongo-driver/mongo"

const (
	mongoPageSpeed = "pagespeed"
)

type Mongo struct {
	client *mongo.Client
}
