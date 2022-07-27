package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
}

func NewClient(ctx context.Context, uri string) (*Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri), options.Client().SetReplicaSet("rs"))
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) Database(name string) *Database {
	return &Database{database: c.client.Database(name)}
}

type Database struct {
	database *mongo.Database
}

func NewDatabase() *Database {
	return new(Database)
}

type Collection struct {
	collection *mongo.Collection
}

func (d *Database) Collection(name string) *Collection {
	return &Collection{collection: d.database.Collection(name)}
}

type Session struct {
	session mongo.Session
}

func (c *Collection) NewTransaction(ctx context.Context, txn func(sc mongo.SessionContext) (interface{}, error)) error {
	session, err := c.collection.Database().Client().StartSession()
	if err != nil {
		return err
	}

	if _, err := session.WithTransaction(ctx, txn); err != nil {
		return err
	}

	return nil
}

func (c *Collection) FindOne(ctx context.Context, filter interface{}, result interface{}) error {
	return c.collection.FindOne(ctx, filter).Decode(result)
}
