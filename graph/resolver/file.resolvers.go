package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/r08521610/personal_backend/graph/model"
	"github.com/r08521610/personal_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	fileID := mongo.GetInstance().SaveFile("ImagesHolder", file.File, file.Filename)
	id := mongo.GetInstance().InsertOne("ImagesHolder", "files", bson.M{
		"name":   file.Filename,
		"fileID": fileID,
	})
	f := model.File{
		ID:     id,
		Name:   file.Filename,
		FileID: fileID,
	}

	return &f, nil
}

func (r *queryResolver) Empty(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
