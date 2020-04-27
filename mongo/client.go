package mongo

import (
	"fmt"
	"os"
	"io"
	"sync"
	"time"
	"bytes"
	"context"

	"github.com/r08521610/personal_backend/graph/model"
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
func (c *Client) InsertOne(databaseName string, collectionName string, b interface{}) string {
	var id string
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	instance.client.Connect(ctx)
	collection := instance.client.Database(databaseName).Collection(collectionName)
	res, _ := collection.InsertOne(ctx, b)

	id = res.InsertedID.(primitive.ObjectID).Hex()
	return id
}

// Find find all data fulfilled filter
func (c *Client) Find(databaseName string, collectionName string, filter bson.M, decode interface{}) []interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	instance.client.Connect(ctx)
	collection := instance.client.Database(databaseName).Collection(collectionName)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		os.Exit(1)
	}
	defer cur.Close(ctx)
	var results []interface{}
	for cur.Next(ctx) {
		result := decode
		err := cur.Decode(result)
		if err != nil {
			os.Exit(1)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		os.Exit(1)
	}
	return results
}

// FindOne get one data
func (c *Client) FindOne(databaseName string, collectionName string, filter bson.M, decode interface{}) interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	instance.client.Connect(ctx)
	collection := instance.client.Database(databaseName).Collection(collectionName)
	err := collection.FindOne(ctx, filter).Decode(decode)
	if err != nil {
		os.Exit(1)
	}
	return decode
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

// DownloadFile download file with gridfs
func (c *Client) DownloadFile(databaseName string, id string) ([]byte, int64) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	instance.client.Connect(ctx)

	collection := instance.client.Database(databaseName).Collection("files")
	fileID, _ := primitive.ObjectIDFromHex(id)
	file := model.File{}
	err := collection.FindOne(ctx, bson.M{"_id": fileID}).Decode(&file)
	if err != nil {
		os.Exit(1)
	}

	bucket, err := gridfs.NewBucket(
		instance.client.Database(databaseName),
	)
	if err != nil {
		os.Exit(1)
	}
	fileID, _ = primitive.ObjectIDFromHex(file.FileID)
	fileBuffer := bytes.NewBuffer(nil)
	size, err := bucket.DownloadToStream(fileID, fileBuffer)
	fmt.Println(err)
	if err != nil {
    os.Exit(1)
	}
	return fileBuffer.Bytes(), size
}
