package mongo

import (
	"os"
	"io"
	"sync"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var once sync.Once
var instance *Client

// Client contains mongo.Client
type Client struct {
	client	*mongo.Client
}

// GetInstance singleton instance
func GetInstance() *Client {
	once.Do(func() {
		c, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		instance = &Client{
			client: c,
		}
	})
	return instance
}

// InsertOne insert data into MongoDB
func (c *Client) InsertOne(databaseName string, collectionName string, b bson.M) string {
	var id string
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	instance.client.Connect(ctx)
	collection := instance.client.Database(databaseName).Collection(collectionName)
	res, _ := collection.InsertOne(ctx, b)

	id = res.InsertedID.(primitive.ObjectID).Hex()
	return id
}

// FindOne get one in bson format
func (c *Client) FindOne(databaseName string, collectionName string, filter bson.M) bson.M {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	instance.client.Connect(ctx)
	collection := instance.client.Database(databaseName).Collection(collectionName)
	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		os.Exit(1)
	}
	return result
}

// SaveFile save file with gridfs
func (c *Client) SaveFile(databaseName string, file io.Reader, filename string) string {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	instance.client.Connect(ctx)
	bucket, err := gridfs.NewBucket(
		instance.client.Database(databaseName),
	)
	if err != nil {
		os.Exit(1)
	}
	fileID, err := bucket.UploadFromStream(
		filename, file,
	)
	if err != nil {
		os.Exit(1)
	}
	return fileID.Hex()
}
