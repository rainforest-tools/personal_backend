package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/r08521610/personal_backend/graph/model"
	"github.com/r08521610/personal_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	id := mongo.GetInstance().InsertOne("ImagesHolder", "projects", bson.M{
		"name": input.Name,
	})
	project := model.Project{
		ID:   id,
		Name: input.Name,
	}
	return &project, nil
}

func (r *mutationResolver) CreateImage(ctx context.Context, input model.NewImage) (*model.Image, error) {
	file := mongo.GetInstance().FindOne("ImagesHolder", "files", bson.M{
		"_id": input.FileID,
	})
	project := mongo.GetInstance().FindOne("ImagesHolder", "projects", bson.M{
		"_id": input.ProjectID,
	})
	id := mongo.GetInstance().InsertOne("ImagesHolder", "images", bson.M{
		"fileID": input.FileID,
		"projectID": input.ProjectID,
	})
	image := model.Image{
		ID: id,
		Project: &model.Project{
			ID: project._id,
			Name: project.name,
		},
		File: &model.File{
			ID: file._id,
			Name: file.name,
			FileID: file.fileID,
		},
	}
	return &image, nil
}

func (r *queryResolver) Images(ctx context.Context) ([]*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}
